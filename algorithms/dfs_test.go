package algorithms

import (
	"log"
	"testing"
)

func TestDFSGraph(t *testing.T) {
	vertexA := GraphVertex{name: "A"}
	vertexB := GraphVertex{name: "B"}
	vertexC := GraphVertex{name: "C"}
	vertexD := GraphVertex{name: "D"}
	vertexE := GraphVertex{name: "E"}
	vertexF := GraphVertex{name: "F"}
	vertexG := GraphVertex{name: "G"}
	vertexH := GraphVertex{name: "H"}
	vertexI := GraphVertex{name: "I"}

	graph := DFSGraph{}
	graph.addGraphVertex(vertexA, GraphVertex{})
	graph.addGraphVertex(vertexB, vertexA)
	graph.addGraphVertex(vertexC, vertexB)
	graph.addGraphVertex(vertexD, vertexB)
	graph.addGraphVertex(vertexE, vertexC)
	graph.addGraphVertex(vertexF, vertexC)
	graph.addGraphVertex(vertexG, vertexD)
	graph.addGraphVertex(vertexH, vertexD)
	graph.addGraphVertex(vertexI, vertexD)

	result := graph.dfsTraverse()
	log.Printf("DFS Result: %v", result)
}
