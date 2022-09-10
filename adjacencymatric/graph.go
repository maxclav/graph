package adjacencymatric

type Graph struct {
	matrix [][]int
}

func NewGraph() *Graph {
	return &Graph{
		matrix: make([][]int, 0),
	}
}
