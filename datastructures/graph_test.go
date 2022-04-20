package datastructures

import (
	"log"
	"testing"
)

func TestAddNodeAndPrintGraph(t *testing.T) {
	log.Printf("Running TestAddNodeAndPrintGraph:::")
	// create graph
	g := Graph{
		vertices: []Vertex{},
		edges:    make(map[Vertex][]Vertex),
	}
	// create nodes
	n1 := Vertex{value: 1}
	n2 := Vertex{value: 2}
	n3 := Vertex{value: 3}

	// add nodes to graph
	g.addVertex(n1)
	g.addVertex(n2)
	g.addVertex(n3)

	// add edges to nodes
	g.addEdge(n1, n2)
	g.addEdge(n1, n3)
	g.addEdge(n2, n3)

	g.printGraph()
}
