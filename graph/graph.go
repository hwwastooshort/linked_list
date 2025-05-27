package graph

type Node[T comparable] struct {
	Value T
}

type Graph[T comparable] struct {
	adjacency map[Node[T]][]Node[T]
}

func NewGraph[T comparable]() *Graph[T] {
	return &Graph[T]{adjacency: make(map[Node[T]][]Node[T])}
}

func (g *Graph[T]) AddNode(value T) Node[T] {

	node := Node[T]{Value: value}

	if _, exists := g.adjacency[node]; !exists {
		g.adjacency[node] = []Node[T]{}
	}

	return node

}

func (g *Graph[T]) AddEdge(from, to Node[T]) {
	g.adjacency[from] = append(g.adjacency[from], to)
}
