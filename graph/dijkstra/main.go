package main

import (
	"container/heap"
	"fmt"
	"math"
	"strconv"

	. "github.com/wwgberlin/algorithms/graph"
	"github.com/wwgberlin/algorithms/graph/pqueue"
)


/*
from wikipedia (Dijkstra's algorithm)
function Dijkstra(Graph, source, target):
	dist[source] ← 0							// Initialization

	create vertex set Q

	for each vertex v in Graph:
		if v ≠ source
			dist[v] ← INFINITY					// Unknown distance from source to v
	  	prev[v] ← nil							// Predecessor of v

	 	Q.add_with_priority(v, dist[v])


	while Q is not empty:						// The main loop
		u ← Q.extract_min()						// Remove and return best vertex
		for each neighbor v of u:				// only v that are still in Q
			alt ← dist[u] + length(u, v)
			if alt < dist[v]
				dist[v] ← alt
				prev[v] ← u
				Q.decrease_priority(v, alt)

	return dist, prev

*/
func Dijkstra(g WGraph, s int, t int) ([]int, int) {
	const Infinity = math.MaxInt32



	//leave this code in at the bottom to return
	//the computed path and the distance

	var out []int
	for v := &t; v != nil; v = prev[*v] {
		out = append([]int{*v}, out...)
	}

	return out, dist[t]
}



func main() {
	g := WGraph(make([][]Edge, 8))

	g = SampleGraphWeighted()

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
		path, distance := Dijkstra(g, c.s, c.t)
		if len(path) == 1 {
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
		fmt.Println(s, fmt.Sprintf("(%d)", distance))
	}
}
