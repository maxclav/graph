package main

import (
	"fmt"
	"math/rand"

	"github.com/maxclav/graph/adjacencylist"
)

const (
	NumberOfVertices int = 10
)

func main() {
	graph := adjacencylist.NewGraph()
	graph.Print()
	fmt.Println()

	fmt.Println("Adding vertex")
	for i := 0; i < NumberOfVertices; i++ {
		vertex := adjacencylist.NewVertex(i)
		graph.AddVertex(vertex)
	}
	graph.Print()
	fmt.Println()

	fmt.Println("Adding random edges")
	for _, from := range graph.Vertices() {
		for _, to := range graph.Vertices() {
			if rand.Intn(4) == 0 {
				from.AddEdgeTo(to)
			}
		}
	}
	graph.Print()
	fmt.Println()
}
