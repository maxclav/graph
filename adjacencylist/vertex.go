package adjacencylist

import "fmt"

type Vertices []*Vertex

type Vertex struct {
	value    int
	adjacent Vertices
}

func NewVertex(val int) *Vertex {
	return &Vertex{
		value: val,
	}
}

func (v *Vertex) Value() int {
	return v.value
}

func (v *Vertex) NumberOfAdjacent() int {
	return len(v.adjacent)
}

func (v *Vertex) HasAdjacent() bool {
	return v.NumberOfAdjacent() > 0
}

func (v *Vertex) Adjacent() Vertices {
	a := make(Vertices, v.NumberOfAdjacent())
	copy(a, v.adjacent)
	return a
}

func (v *Vertex) AddEdgeTo(a *Vertex) {
	v.adjacent = append(v.adjacent, a)
}

func (v *Vertex) RemoveEdgeOf(a *Vertex) {
	for idx, item := range v.adjacent {
		if item == a {
			v.adjacent = append(v.adjacent[:idx], v.adjacent[idx+1:]...)
		}
	}
}

func (v *Vertex) WithEdgeTo(a *Vertex) *Vertex {
	v.AddEdgeTo(a)
	return v
}

func (v *Vertex) WithoutEdgeTo(a *Vertex) *Vertex {
	v.RemoveEdgeOf(a)
	return v
}

func (v *Vertex) Print() {
	fmt.Printf("[%v] -> ", v.value)
	if !v.HasAdjacent() {
		fmt.Print("none!")
	} else {
		for _, a := range v.adjacent[:v.NumberOfAdjacent()-1] {
			fmt.Printf("%v, ", a.Value())
		}
		fmt.Printf("%v", v.adjacent[v.NumberOfAdjacent()-1].Value())
	}
}
