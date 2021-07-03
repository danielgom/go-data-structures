package main

import (
	"fmt"
)

// Graph structure represents and adjacency list graph
type Graph struct {
	vertices []*Vertex
}

// Vertex structure represents a graph vertex
type Vertex struct {
	key      int
	adjacent []*Vertex
}

// AddVertex adds a vertex to the graph
func (g *Graph) AddVertex(k int) {
	if contains(g.vertices, k) {
		err := fmt.Errorf("vertex %v was not added because it is and existing key", k)
		fmt.Println(err.Error())
	} else {
		g.vertices = append(g.vertices, &Vertex{key: k})
	}
}

// AddEdge adds an edge to the graph (directed, means just add to one direction)
func (g *Graph) AddEdge(from, to int) {
	// get vertex
	fromVertex := g.getVertex(from)
	toVertex := g.getVertex(to)
	// check error
	if fromVertex == nil || toVertex == nil {
		err := fmt.Errorf("invalid edge (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else if contains(fromVertex.adjacent, to) {
		err := fmt.Errorf("edge already exists (%v --> %v)", from, to)
		fmt.Println(err.Error())
	} else {
		// add edge
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// getVertex
func (g *Graph) getVertex(k int) *Vertex {
	for i, v := range g.vertices {
		if k == v.key {
			return g.vertices[i]
		}
	}
	return nil
}

// contains
func contains(s []*Vertex, k int) bool {
	for _, v := range s {
		if k == v.key {
			return true
		}
	}
	return false
}

// Print will print the adjacent list for each vertex of the graph
func (g *Graph) Print() {
	for _, v := range g.vertices {
		fmt.Printf("\nVertex %d:", v.key)
		for _, vs := range v.adjacent {
			fmt.Print(" ", vs.key)
		}
	}
	fmt.Println()
}

func main() {

	tg := &Graph{}

	for x := 0; x < 5; x++ {
		tg.AddVertex(x)
	}

	tg.Print()
	tg.AddVertex(5)
	tg.Print()

	tg.AddEdge(1, 2)
	tg.AddEdge(1, 2)
	tg.AddEdge(6, 7)
	tg.Print()

}
