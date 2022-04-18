package algorithms

import (
	"log"
	"testing"
)

func TestAddNodeAndPrintGraph(t *testing.T) {
	log.Printf("Running TestAddNodeAndPrintGraph:::")
	// create graph
	g := Graph{
		nodes: []Node{},
		edges: make(map[Node][]Node),
	}
	// create nodes
	n1 := Node{value: 1}
	n2 := Node{value: 2}
	n3 := Node{value: 3}

	// add nodes to graph
	g.addNode(n1)
	g.addNode(n2)
	g.addNode(n3)

	// add edges to nodes
	g.addEdge(n1, n2)
	g.addEdge(n1, n3)
	g.addEdge(n2, n3)

	g.printGraph()
}
