package algorithms

import "log"

type GraphVertex struct {
	name    string
	visited bool
}

type GraphEdge struct {
	edges map[GraphVertex][]GraphVertex
}

type DFSGraph struct {
	vertices []GraphVertex
	adjLists map[GraphVertex][]GraphVertex
}

func (dfsGrgaph *DFSGraph) addGraphVertex(vertex GraphVertex, to GraphVertex) {
	if dfsGrgaph.vertices == nil {
		dfsGrgaph.vertices = []GraphVertex{}
	}
	if dfsGrgaph.adjLists == nil {
		dfsGrgaph.adjLists = make(map[GraphVertex][]GraphVertex)
	}
	dfsGrgaph.vertices = append(dfsGrgaph.vertices, vertex)
	toVertices := dfsGrgaph.adjLists[vertex]
	if toVertices == nil {
		toVertices = []GraphVertex{}
	}
	toVertices = append(toVertices, to)
	if (to != GraphVertex{}) {
		dfsGrgaph.adjLists[vertex] = toVertices
	}
	log.Printf("Addeded the vertex %v to %v", vertex, to)
	log.Printf("Vertices are: %v with edges: %v", dfsGrgaph.vertices, dfsGrgaph.adjLists)
}

func (dfsGraph *DFSGraph) dfsTraverse() []GraphVertex {
	currVertices := dfsGraph.vertices
	adjLists := dfsGraph.adjLists
	visited_vertices := []GraphVertex{}
	return dfsTraverseHelper(currVertices, &adjLists, visited_vertices)
}

func dfsTraverseHelper(vertices []GraphVertex, adjLists *map[GraphVertex][]GraphVertex, visited_vertices []GraphVertex) []GraphVertex {
	if vertices == nil || len(vertices) == 0 {
		return []GraphVertex{}
	}
	for i, vertex := range vertices {
		if visited(visited_vertices, vertex) {
			log.Printf("Vertex: %v visited and so skipping", vertex)
			continue
		}
		// if vertex.visited {
		// 	log.Printf("Vertex: %v visited and so skipping", vertex)
		// 	continue
		// }
		log.Printf("ith: %d, Vertex: %v", i, vertex)
		vertices = append(vertices, vertex)
		connected_vertices := (*adjLists)[vertex]
		vertex.visited = true
		visited_vertices = append(visited_vertices, vertex)
		log.Printf("Connected vertices: %v", connected_vertices)
		dfsTraverseHelper(connected_vertices, adjLists, vertices)
	}
	return vertices
}

func visited(s []GraphVertex, e GraphVertex) bool {
	log.Printf("Visited check::: %v", s)
	log.Printf("Vertex::: %v", e)
	for _, a := range s {
		if a.name == e.name && a.visited {
			log.Printf("Already visited::: %v", e)
			// if a == e {
			return true
		}
	}
	return false
}
