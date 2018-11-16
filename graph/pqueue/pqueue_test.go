package pqueue

import (
	"container/heap"
	"testing"

	. "github.com/wwgberlin/algorithms/graph"
)

func TestPQueue_Swap(t *testing.T) {
	pq := NewPQueue()
	for i := 0; i < 5; i++ {
		pq.AddWithPriority(i, 1000)
	}
	pq.pairs[3].D = 1
	pq.pairs[4].D = 2

	pq.Swap(3, 4)

	if pq.positions[4] != 3 {
		t.Error("Unexpected result after Swap", pq.positions[4])
	}

	if pq.pairs[4].D != 1 {
		t.Error("Unexpected result after Swap", pq.pairs[4])
	}

	if pq.positions[3] != 4 {
		t.Error("Unexpected result after Swap", pq.positions[3])
	}
	if pq.pairs[3].D != 2 {
		t.Error("Unexpected result after Swap", pq.pairs[3])
	}

}

func TestPQueue_Less(t *testing.T) {
	pq := NewPQueue()
	for i := 0; i < 5; i++ {
		pq.AddWithPriority(i, 1000)
	}
	pq.pairs[3].D = 1
	pq.pairs[4].D = 2
	if !pq.Less(3, 4) {
		t.Error("1<2 expected to return true")
	}
	if pq.Less(4, 3) {
		t.Error("Less is inconsistent")
	}

}

func TestPQueue_DecreaseKey(t *testing.T) {
	pq := NewPQueue()
	for i := 0; i < 5; i++ {
		pq.AddWithPriority(i, 1000)
	}
	heap.Init(pq)

	pq.DecreaseKey(4, 0)

	e := pq.pairs[0]
	if e.V != 4 || e.D != 0 {
		t.Error("Unexpected result after DecreaseKey", pq.pairs)
	}
	if pq.positions[4] != 0 {
		t.Error("Lost track of item after DecreaseKey", pq.positions)
	}

	pq.DecreaseKey(3, 5)

	e = pq.pairs[1]
	if e.V != 3 || e.D != 5 {
		t.Error("Unexpected result after DecreaseKey", pq.pairs)
	}
	if pq.positions[3] != 1 {
		t.Error("Lost track of item after DecreaseKey", pq.positions)
	}
}

func TestPQueue_Pop(t *testing.T) {
	pq := NewPQueue()
	for i := 0; i < 4; i++ {
		pq.AddWithPriority(i, 1000)
	}
	for i := range pq.pairs {
		pq.pairs[i].D = i
	}
	for i := range pq.pairs {
		v := heap.Pop(pq).(int)
		if v != i {
			t.Error("Unexpected item popped", v, "at iteration", i)
		}
		if _, ok := pq.positions[i]; ok {
			t.Error("Position of vertex was not deleted")
		}
	}

}

func TestPQueue_Push(t *testing.T) {
	//0 has been popped and can't be pushed again
	pq := PQueue{
		pairs:     []Edge{{V: 1, D: 3}, {V: 2, D: 4}},
		positions: map[int]int{1: 0, 2: 1},
	}

	heap.Push(&pq, Edge{0, -10})
	if pq.positions[0] != 0 {
		t.Error("Unexpected position for item")
	}
	if pq.pairs[0].V != 0 || pq.pairs[0].D != -10 {
		t.Error("Unexpected min item")
	}

}
