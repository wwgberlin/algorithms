package main

import (
	"math/rand"
	"time"
	"fmt"
	"container/heap"
	"math"
	. "github.com/wwgberlin/algorithms/graph"
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
	fmt.Println(Dijkstra(g, 2, 7))
	fmt.Println(Dijkstra(g, 1, 7))
	fmt.Println(Dijkstra(g, 3, 4))
	fmt.Println(Dijkstra(g, 4, 2))
	fmt.Println(Dijkstra(g, 4, 3))
	fmt.Println(Dijkstra(g, 5, 1))
	fmt.Println(Dijkstra(g, 5, 7))
	fmt.Println(Dijkstra(g, 5, 3))
	fmt.Println(Dijkstra(g, 6, 3))
	fmt.Println(Dijkstra(g, 6, 2))
}

/*
Dijkstra(g,s) //simplified
	pq := make empty priority queue
	dist : make empty map from vertices to distances (weights)
	prev := make empty map from vertex to vertex

	pq.push(s, 0) //the distance of s from the source is 0
	dist[src] := 0 //initialize the correct distance for the vertex s

	while pq is not empty
			u := pq.pop
			pq.pop();
			for each v neighbor of u
		   alt := dist[u] + weight
				if (dist[v] > alt) then
					dist[v] := alt
			   prev[v]
					pq.push(v, alt);


	// output the shortest path to t:
		out:= make empty slice
		v,ok:= t, prev[t]
		while ok
			out:=prepend(v, out)
			v,ok:= prev[v]

		return out, dist[t]
 */
func Dijkstra(g Graph, s int, t int) ([]int, int) {
	pq := PQueue(make([]Edge, 0))
	heap.Push(&pq, Edge{s, 0})

	dist := make(map[int]int)

	for i := range g {
		dist[i] = math.MaxInt32
	}

	prev := make(map[int]int)
	dist[s] = 0

	for len(pq) > 0 {
		e := heap.Pop(&pq).(Edge)
		u := e.V

		for _, e := range g[e.V] {
			v := e.V
			d := e.D

			if dist[v] > dist[u]+d {
				dist[v] = dist[u] + d
				prev[v] = u
				heap.Push(&pq, Edge{v, dist[v],})
			}
		}
	}

	var out []int

	for v, ok := prev[t]; ok; v, ok = prev[v] {
		out = append([]int{v}, out...)
		if v == s {
			break
		}
	}

	return out, dist[t]
}
