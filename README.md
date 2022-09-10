# Graph in Go

Simple directed graph where vertices have int as values and edges are not weighted.  

Implementation with:
- Adjacency list: `not done`
- Adjacency matrix: `not started`

For learning purposes only.  

## Commands

- `make` or `make help`: display all commands
- `make run-lint`: run linter
- `make run-tests`: unit tests not completed
- `make example`: run example

## TODO/Ideas

- Add unit tests
- Add unit interface/functions/methods (is directed/undirected, is cyclic, ...)
- Change slice of vertices (`type Vertices []*Vertex`) to a set of vertices (`type Certices map[*Vertex]struct{}`) so we don't have duplicated vertices/edges in the graph
