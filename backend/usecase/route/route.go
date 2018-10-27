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

	var cs []master.StationMaster
	for _, v := range u.cuc.GetStations() {
		if verify(v) {
			cs = append(cs, v)
		}
	}

	g := NewGraph()
	for _, v := range cs {
		for _, vs := range v.Railway.StationOrders {
			if exists(cs, vs.StationSameAs) {
				fare, err := u.cuc.GetRailwayFare(v.SameAs, vs.StationSameAs)
				if err != nil {
					continue
				}
				g.Add(v.SameAs, vs.StationSameAs, fare.IcCardFare)
			}
		}
		for _, r := range v.ConnectingRailways {
			if r.Railway == nil {
				continue
			}
			for _, rs := range r.Railway.StationOrders {
				if exists(cs, rs.StationSameAs) {
					fare, err := u.cuc.GetRailwayFare(v.SameAs, rs.StationSameAs)
					if err != nil {
						continue
					}
					g.Add(v.SameAs, rs.StationSameAs, fare.IcCardFare)
				}
			}
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

func exists(cs []master.StationMaster, stationSameAs string) bool {
	for _, v := range cs {
		if v.SameAs == stationSameAs {
			return true
		}
	}
	return false
}

const NIL = -1

type Graph struct {
	Nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{map[string]*Node{}}
}

type Tree struct {
	Next *Node
	Cost int
}

type Node struct {
	Name  string
	Trees []*Tree
	Done  bool
	Cost  int
	Prev  *Node
}

func NewNode(name string) *Node {
	return &Node{
		Name:  name,
		Trees: []*Tree{},
		Done:  false,
		Cost:  NIL,
		Prev:  nil,
	}
}

func NewTree(next *Node, cost int) *Tree {
	return &Tree{Next: next, Cost: cost}
}

func (n *Node) AddTree(tree *Tree) {
	n.Trees = append(n.Trees, tree)
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
	tree := NewTree(toNode, cost)
	fromNode.AddTree(tree)
}

func (g *Graph) Route(start string, goal string) (ret []*Node, err error) {
	startNode := g.Nodes[start]
	startNode.Cost = 0

	for {
		node, err := g.nextNode()
		if err != nil {
			return nil, errors.New("Goal not found")
		}
		if node.Name == goal {
			break
		}
		g.calculate(node)
	}

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

func (g *Graph) calculate(node *Node) {
	for _, tree := range node.Trees {
		nextNode := tree.Next
		if nextNode.Done {
			continue
		}
		cost := node.Cost + tree.Cost
		if nextNode.Cost == NIL || cost < nextNode.Cost {
			nextNode.Cost = cost
			nextNode.Prev = node
		}
	}
	node.Done = true
}

func (g *Graph) nextNode() (next *Node, err error) {
	for _, node := range g.Nodes {
		if node.Done {
			continue
		}
		if node.Cost == -1 {
			continue
		}
		if next == nil {
			next = node
		}
		if next.Cost > node.Cost {
			next = node
		}
	}
	if next == nil {
		return nil, errors.New("Untreated node not found")
	}
	return
}
