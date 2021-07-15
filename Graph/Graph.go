package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"unicode/utf8"
)

type Reduction []int

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

	res, err := http.Get("https://www.google.com")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(body))

	f := Filter([]int{1, 2, 3, 4}, func(b int) bool {
		return b%2 == 0
	})

	fmt.Println(f)

	fil := func(b int) bool {
		return b%2 == 0
	}

	f2 := Filter([]int{1, 2, 3, 4}, fil)
	fmt.Println(f2.Reduce())
	fmt.Println()

	f3 := []int{1, 2, 3, 4}

	sort.Sort(sort.Reverse(sort.IntSlice(f3)))

	b := []byte{'z', 'b'}
	st := "asd"

	decodeRune, size := utf8.DecodeRune(b)
	fmt.Println(decodeRune)
	fmt.Println(size)
	fmt.Println(utf8.RuneCountInString(st))

	characterReps := make(map[byte]int)
	characterReps['a'] = 0
	characterReps['b'] = 1
	characterReps['c'] = 2

	for i, v := range characterReps {
		fmt.Println(i)
		fmt.Println(v)
	}

	i := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println(i[len(i)/2:])

	fmt.Println(create())

	arr := create()

	arr[0] = 22
	fmt.Println(arr[0])
	fmt.Println(arr)

	is := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	is = append(is[:2], is[3:]...)
	fmt.Println(is)

}

func Filter(sl []int, a func(b int) bool) Reduction {
	res := make([]int, 0)

	for _, v := range sl {
		if a(v) {
			res = append(res, v)
		}
	}
	return res
}

func (r *Reduction) Reduce() int {
	res := 0
	for _, v := range *r {
		res += v
	}
	return res
}

func create() []int {
	return make([]int, 10)
}
