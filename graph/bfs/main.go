package main

import (
	. "github.com/wwgberlin/algorithms/graph"
	"fmt"
	"math"
	"strconv"
)

const (
	_         = iota
	White
	Gray
	Black
	Inifinity = math.MaxInt32
)

/*
From Introduction to Algorithms (Corman):
BFS(G, s)
	for each vertex u in G.V - {s}
 		u.color = WHITE
 		u.d = INFIINITE
 		u.pi= nil
	s.color = GRAY
	s.d = 0
	s.pi = nil
	Q = {}
	ENQUEUE(Q, s)
	while Q not empty
		u = DEQUEUE(Q)
		for each v in G.Adj[u]
			if v.color == WHITE
				v.color = GRAY
				v.d = u.d + 1
				v.pi = u
				ENQUEUE(Q,v)
			u.color = BLACK
 */
func BFS(g Graph, s int, t int) ([]int, int) {
	queue := []int{s}
	pi := make([]int, len(g))
	dist := make([]int, len(g))
	color := make([]uint8, len(g))

	for i := range g {
		dist[i] = Inifinity
		color[i] = White
		pi[i] = -1
	}

	dist[s] = 0
	color[s] = Gray
	pi[s] = -1

	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		for _, v := range g[u] {
			if color[v] == White {
				color[v] = Gray
				dist[v] = dist[u] + 1
				pi[v] = u
				queue = append(queue, v)
			}
			color[u] = Black
		}
	}

	var out []int
	for v := t; v != -1; v = pi[v] {
		out = append([]int{v}, out...)
	}
	return out, dist[t]
}

func main() {
	g := SampleGraph()
	fmt.Println(g)

	cases := []struct {
		s int
		t int
	}{
		{2, 7},
		{1, 7},
		{3, 4},
		{4, 2},
		{4, 3},
		{5, 1},
		{5, 7},
		{5, 3},
		{6, 3},
		{6, 2},
		{7, 1},
	}

	for _, c := range cases {
		fmt.Printf("Calculating shortest path from %d to %d:\n", c.s, c.t)
		path, distance := BFS(g, c.s, c.t)
		if distance == math.MaxInt32 {
			fmt.Println("No path found")
			continue
		}
		var s string
		for i, v := range path {
			if i != 0 {
				s += "->"
			}
			s += strconv.Itoa(v)
		}
		fmt.Println(s, fmt.Sprintf("(Distance: %d)", distance))
	}
}
