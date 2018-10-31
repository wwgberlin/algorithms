package main

import (
	"math/rand"
	"time"
	"fmt"
	"container/heap"
	"math"
	. "github.com/wwgberlin/algorithms/graph"
)

type PQueue struct {
	arr  []int
	dist map[int]int
}

func (q PQueue) Len() int {
	return len(q.arr)
}
func (q PQueue) Less(i, j int) bool {
	return q.dist[q.arr[i]] < q.dist[q.arr[j]]
}

func (q PQueue) Swap(i, j int) {
	q.arr[i], q.arr[j] = q.arr[j], q.arr[i]
}

func (q *PQueue) Push(x interface{}) {
	q.arr = append(q.arr, x.(int))
}

func (q *PQueue) Pop() interface{} {
	old := q.arr
	n := len(old)
	x := old[n-1]
	q.arr = old[0: n-1]
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

func Dijkstra(g Graph, s int, t int) ([]int, int) {

	pq := PQueue{make([]int, 0), make(map[int]int)}
	pq.dist[s] = 0

	for v := range g {
		if v != s {
			pq.dist[v] = math.MaxInt32
		}
		heap.Push(&pq, v)
	}

	prev := make(map[int]int)

	visited := make([]bool, len(g))

	for pq.Len() > 0 {
		u := heap.Pop(&pq).(int)
		fmt.Println("visiting:", u)
		visited[u] = true

		for _, e := range g[u] {
			v := e.V

			if visited[v] {
				continue
			}

			d := e.D
			alt := pq.dist[u] + d

			if alt < pq.dist[v] {
				pq.dist[v] = alt
				heap.Init(&pq)
				prev[v] = u
			}
		}
	}

	var out []int

	for v, ok := t, true; ok; v, ok = prev[v] {
		out = append([]int{v}, out...)
		if v == s {
			break
		}
	}

	if d, ok := pq.dist[t]; ok {
		return out, d
	}
	return out, math.MaxInt32
}
