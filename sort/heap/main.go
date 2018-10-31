package main

import (
	"container/heap"
	"fmt"
)

func main() {
	s := []int{5, 2, 3, 5, 9, 1, 2, 2}
	heapSort(s)
	fmt.Println(s)
}
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

//Start by defining a new duplicate slice dup that points to the same backing
//array as given slice a (no allocation necessary)
//Define a new variable h of type IntHeap that points to dup.
//Use heap.Init(&h) to initialize a heap from the slice.
//Note that we have to send the pointer to h.
//For each index in a, use heap.Pop(&h) to retrieve the first element in the
//heap, and put it at the relative index of the item at the back of the slice
//that would be len(a)-i-1
//every popped item will have to be converted back to an integer as Pop returns
//the empty interface (interface{})
func heapSort(a []int) {
	dup := a
	h := IntHeap(dup)

	heap.Init(&h)

	for i := range a {
		a[len(a)-i-1] = heap.Pop(&h).(int)
	}
}
