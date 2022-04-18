package algorithms

import (
	"fmt"
	"log"
	"sync"
)

type Node struct {
	value int
}

type Graph struct {
	nodes []Node
	edges map[Node][]Node
	lock  sync.RWMutex
}

func (n Node) String() string {
	return fmt.Sprintf("%v:", n.value)
}

func (g Graph) printGraph() {
	g.lock.Lock()
	log.Printf("Printing Graph:")
	nodes := g.nodes
	for _, n := range nodes {
		conn_nodes := g.edges[n]
		log.Printf("Node: %v\n", n.String())
		for _, cn := range conn_nodes {
			log.Printf("\tNode %v\n", cn.String())
		}
	}
	g.lock.Unlock()
}

func (g *Graph) addNode(n Node) {
	log.Printf("Adding node: %v", n.String())
	g.lock.Lock()
	g.nodes = append(g.nodes, n)
	g.lock.Unlock()
}

func (g *Graph) addEdge(from Node, to Node) {
	log.Printf("Adding edge: from %v to %v", from.String(), to.String())
	nodes := append(g.edges[from], to)
	g.edges[from] = nodes
}
