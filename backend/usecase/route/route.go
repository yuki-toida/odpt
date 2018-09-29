package route

import (
	"errors"
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
		if 0 < len(v.ConnectingRailways) {
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
	fare, err := u.cuc.GetRailwayFare(from.SameAs, to.SameAs)
	if err != nil {
		return nil, errors.New("運賃を取得できませんでした")
	}

	res := &struct{ From, To, Fare interface{} }{From: from, To: to, Fare: fare}
	return res, nil
}

const NIL = -1

type DirectedGraph struct {
	nodes map[string]*Node
}

func NewDirectedGraph() *DirectedGraph {
	return &DirectedGraph{map[string]*Node{}}
}

type Edge struct {
	next *Node
	cost int
}

type Node struct {
	name  string
	edges []*Edge
	done  bool
	cost  int
	prev  *Node
}

func NewNode(name string) *Node {
	return &Node{
		name:  name,
		edges: []*Edge{},
		done:  false,
		cost:  NIL,
		prev:  nil,
	}
}

func NewEdge(next *Node, cost int) *Edge {
	return &Edge{next: next, cost: cost}
}

func (n *Node) AddEdge(edge *Edge) {
	n.edges = append(n.edges, edge)
}

func (dg *DirectedGraph) Add(from, to string, cost int) {
	fromNode, ok := dg.nodes[from]
	if !ok {
		fromNode = NewNode(from)
		dg.nodes[from] = fromNode
	}

	toNode, ok := dg.nodes[to]
	if !ok {
		toNode = NewNode(to)
		dg.nodes[to] = toNode
	}

	edge := NewEdge(toNode, cost)
	fromNode.AddEdge(edge)
}
