package datastructures

import (
	"fmt"
	"log"
	"sync"
)

type Vertex struct {
	value int
}

type Graph struct {
	vertices []Vertex
	edges    map[Vertex][]Vertex
	lock     sync.RWMutex
}

func (n Vertex) String() string {
	return fmt.Sprintf("%v:", n.value)
}

func (g Graph) printGraph() {
	g.lock.Lock()
	log.Printf("Printing Graph:")
	vertices := g.vertices
	for _, n := range vertices {
		conn_nodes := g.edges[n]
		log.Printf("Vertex: %v\n", n.String())
		for _, cn := range conn_nodes {
			log.Printf("\tVertex %v\n", cn.String())
		}
	}
	g.lock.Unlock()
}

func (g *Graph) addVertex(n Vertex) {
	log.Printf("Adding vertex: %v", n.String())
	g.lock.Lock()
	g.vertices = append(g.vertices, n)
	g.lock.Unlock()
}

func (g *Graph) addEdge(from Vertex, to Vertex) {
	log.Printf("Adding edge: from %v to %v", from.String(), to.String())
	vertices := append(g.edges[from], to)
	g.edges[from] = vertices
}
