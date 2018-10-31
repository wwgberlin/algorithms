package main

import (
	"math/rand"
	"time"
	"fmt"
	"container/heap"
	"math"
	. "github.com/wwgberlin/algorithms/graph"
	"strconv"
)

type PQueue []Edge

func (pq PQueue) Len() int {
	return len(pq)
}
func (pq PQueue) Less(i, j int) bool {
	return pq[i].D < pq[j].D
}

func (pq PQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Edge))
}

func (pq *PQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0: n-1]
	return x
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	const l = 5

	g := Graph(make([][]Edge, 8))

	g = SampleGraph()

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
		if len(path) == 0 {
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

/*
Dijkstra(g,s) //simplified
	pq := make empty priority queue
	dist : make empty map from vertices to distances (weights)

	for every vertex v in graph
		dist[v] := infinite (math.MaxInt32)
	prev := make empty map from vertex to vertex

	pq.push(s, 0) //the distance of s from the source is 0
	dist[src] := 0 //initialize the correct distance for the vertex s

	while pq is not empty
		u := pq.pop
		for each v neighbor of u
			alt := dist[u] + weight
			if (dist[v] > alt) then
				dist[v] := alt
			   	prev[v] := u
				pq.push(v, alt);

		//almost done :) output the shortest path to vertex t
		out := make empty slice
		v,ok := t, prev[t]
		while ok
			out:=prepend(v, out)
			v,ok:= prev[v]

		return out, dist[t]
 */
func Dijkstra(g Graph, s int, t int) ([]int, int) {
}
