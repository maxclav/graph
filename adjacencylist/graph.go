package adjacencylist

import "fmt"

type Graph struct {
	vertices Vertices
}

func NewGraph() *Graph {
	return &Graph{
		vertices: make(Vertices, 0),
	}
}

func (g *Graph) Vertices() Vertices {
	v := make(Vertices, g.NumberOfVertices())
	copy(v, g.vertices)
	return v
}

func (g *Graph) AddVertex(v *Vertex) {
	g.vertices = append(g.vertices, v)
}

func (g *Graph) RemoveVertex(v *Vertex) {
	var index int = -1
	for i, vertex := range g.vertices {
		if vertex == v {
			index = i
		}
		vertex.RemoveEdgeOf(vertex)
	}
	if index >= 0 {
		g.vertices = append(g.vertices[:index], g.vertices[index+1:]...)
	}
}

func (g *Graph) WithVertex(v *Vertex) *Graph {
	g.AddVertex(v)
	return g
}

func (g *Graph) WithoutVertex(v *Vertex) *Graph {
	g.RemoveVertex(v)
	return g
}

func (g *Graph) AddEdge(from, to *Vertex) {
	for _, vertex := range g.vertices {
		if vertex == from {
			vertex.AddEdgeTo(to)
		}
	}
}

func (g *Graph) RemoveEdge(from, to *Vertex) {
	for _, vertex := range g.vertices {
		if vertex == from {
			vertex.RemoveEdgeOf(to)
		}
	}
}

func (g *Graph) WithEdge(from, to *Vertex) *Graph {
	g.AddEdge(from, to)
	return g
}

func (g *Graph) WithoutEdge(from, to *Vertex) *Graph {
	g.RemoveEdge(from, to)
	return g
}

func (g *Graph) NumberOfVertices() int {
	return len(g.vertices)
}

func (g *Graph) HasVertices() bool {
	return g.NumberOfVertices() > 0
}

func (g *Graph) Print() {
	fmt.Print("Graph: ")
	if !g.HasVertices() {
		fmt.Println("empty!")
	} else {
		fmt.Println()
		for _, vertex := range g.vertices {
			vertex.Print()
			fmt.Println()
		}
	}
}
