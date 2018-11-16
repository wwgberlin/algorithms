package pqueue

import (
	"container/heap"

	. "github.com/wwgberlin/algorithms/graph"
)

type PQueue struct {
	pairs []Edge //the heap itself
	//using a map here to maintain the dijkstra interface of adding
	//items to the queue dynamically, this can be traded with a slice of a given size
	positions map[int]int //track the positions of the vertices as they are swapped
}

func NewPQueue() *PQueue {
	pairs := make([]Edge, 0)
	positions := make(map[int]int)
	return &PQueue{pairs: pairs, positions: positions}
}

func (pq *PQueue) AddWithPriority(v int, dist int) {
	heap.Push(pq, Edge{v, dist})
}

func (pq PQueue) Len() int {
	return len(pq.pairs)
}
func (pq PQueue) Less(i, j int) bool {
	return pq.pairs[i].D < pq.pairs[j].D
}

func (pq PQueue) Swap(i, j int) {
	pq.pairs[i], pq.pairs[j] = pq.pairs[j], pq.pairs[i]
	pq.positions[pq.pairs[i].V] = i
	pq.positions[pq.pairs[j].V] = j
}

func (pq *PQueue) Push(x interface{}) {
	n := len(pq.pairs)
	e := x.(Edge)
	if _, ok := pq.positions[e.V]; ok {
		panic("trying to push already existing edge to priority queue")
	}
	pq.positions[e.V] = n
	pq.pairs = append(pq.pairs, e)
}

func (pq *PQueue) DecreaseKey(k int, value int) {
	pq.pairs[pq.positions[k]].D = value
	heap.Fix(pq, pq.positions[k])
}

func (pq *PQueue) Pop() interface{} {
	old := pq.pairs
	n := len(old)
	x := old[n-1]
	pq.pairs = old[0: n-1]
	delete(pq.positions, x.V)
	return x.V
}
