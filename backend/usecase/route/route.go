package route

import (
	"errors"
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
	"github.com/yuki-toida/refodpt/backend/domain/model/master"
	"github.com/yuki-toida/refodpt/backend/infrastructure/repository"
	"github.com/yuki-toida/refodpt/backend/usecase/cache"
)

type UseCase struct {
	db  *gorm.DB
	cuc *cache.UseCase
}

func NewUseCase(r *repository.Repository) *UseCase {
	return &UseCase{
		db:  r.DB,
		cuc: cache.NewUseCase(r.Cache),
	}
}

func (u *UseCase) GetConnectingStations() []master.StationMaster {
	var stations []master.StationMaster
	for _, v := range u.cuc.GetStations() {
		if verify(v) {
			stations = append(stations, v)
		}
	}
	sort.Slice(stations, func(i, j int) bool {
		return stations[i].Railway.SameAs < stations[j].Railway.SameAs
	})
	return stations
}

func (u *UseCase) GetRoutes(fromStation, toStation string) (interface{}, error) {
	from, err := u.cuc.GetStationByName(fromStation)
	if err != nil {
		return nil, errors.New("出発駅名が存在しません")
	}
	to, err := u.cuc.GetStationByName(toStation)
	if err != nil {
		return nil, errors.New("到着駅名が存在しません")
	}

	g := NewGraph()
	for _, v := range u.cuc.GetRailwayFares() {
		sf, _ := u.cuc.GetStation(v.FromStationSameAs)
		st, _ := u.cuc.GetStation(v.ToStationSameAs)
		if verify(sf) && verify(st) {
			g.Add(v.FromStationSameAs, v.ToStationSameAs, v.IcCardFare)
		}
	}

	nodes, err := g.Route(from.SameAs, to.SameAs)
	if err != nil {
		return nil, err
	}

	res := make([]struct {
		Name string
		Cost int
	}, len(nodes))

	for i, node := range nodes {
		res[i].Name = node.Name
		res[i].Cost = node.Cost
		fmt.Printf("%v: %v\n", node.Name, node.Cost)
	}

	return res, nil
}

func verify(s master.StationMaster) bool {
	return s.OperatorSameAs == "odpt.Operator:TokyoMetro" && 0 < len(s.ConnectingRailways)
}

const NIL = -1

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{map[string]*Node{}}
}

type Edge struct {
	Next *Node
	Cost int
}

type Node struct {
	Name  string
	Edges []*Edge
	Done  bool
	Cost  int
	Prev  *Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Edges: []*Edge{},
		Done:  false,
		Cost:  NIL,
		Prev:  nil,
	}
}

func NewEdge(next *Node, cost int) *Edge {
	return &Edge{Next: next, Cost: cost}
}

func (n *Node) AddEdge(edge *Edge) {
	n.Edges = append(n.Edges, edge)
}

func (g *Graph) Add(from, to string, cost int) {
	fromNode, ok := g.Nodes[from]
	if !ok {
		fromNode = NewNode(from)
		g.Nodes[from] = fromNode
	}

	toNode, ok := g.Nodes[to]
	if !ok {
		toNode = NewNode(to)
		g.Nodes[to] = toNode
	}

	edge := NewEdge(toNode, cost)
	fromNode.AddEdge(edge)
}

func (g *Graph) Route(start string, goal string) (ret []*Node, err error) {
	// 名前からスタート地点のノードを取得する
	startNode := g.Nodes[start]

	// スタートのコストを 0 に設定することで処理対象にする
	startNode.Cost = 0

	for {
		// 次の処理対象のノードを取得する
		node, err := g.nextNode()

		// 次に処理するノードが見つからなければ終了
		if err != nil {
			return nil, errors.New("Goal not found")
		}

		// ゴールまで到達した
		if node.Name == goal {
			break
		}

		// 取得したノードを処理する
		g.calc(node)
	}

	// ゴールから逆順にスタートまでノードをたどっていく
	n := g.Nodes[goal]
	for {
		ret = append(ret, n)
		if n.Name == start {
			break
		}
		n = n.Prev
	}

	return ret, nil
}

func (g *Graph) calc(node *Node) {
	// ノードにつながっているエッジを取得する
	for _, edge := range node.Edges {
		nextNode := edge.Next

		// 既に処理済みのノードならスキップする
		if nextNode.Done {
			continue
		}

		// このノードに到達するのに必要なコストを計算する
		cost := node.Cost + edge.Cost
		if nextNode.Cost == NIL || cost < nextNode.Cost {
			// 既に見つかっている経路よりもコストが小さければ処理中のノードを遷移元として記録する
			nextNode.Cost = cost
			nextNode.Prev = node
		}
	}

	// つながっているノードのコスト計算がおわったらこのノードは処理済みをマークする
	node.Done = true
}

func (g *Graph) nextNode() (next *Node, err error) {
	for _, node := range g.Nodes {
		// 処理済みのノードは対象外
		if node.Done {
			continue
		}
		// コストが初期値 (-1) になっているノードはまだそのノードまでの最短経路が判明していないので処理できない
		if node.Cost == -1 {
			continue
		}
		// 最初に見つかったものは問答無用で次の処理対象の候補になる
		if next == nil {
			next = node
		}
		// 既に見つかったノードよりもコストの小さいものがあればそちらを先に処理しなければいけない
		if next.Cost > node.Cost {
			next = node
		}
	}
	// 次の処理対象となるノードが見つからなかったときはエラー
	if next == nil {
		return nil, errors.New("Untreated node not found")
	}
	return
}
