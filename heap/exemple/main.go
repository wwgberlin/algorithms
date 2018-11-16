package main

import (
	"container/heap"
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"
)

//Copied from the IntHeap example in container/heap docs
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0: n-1]
	return x
}

func MaxElements(s []int, m int) []int {
	if m > len(s) {
		panic("m must be smaller than the length of slice s")
	}

	out := make([]int, m)
	//O(m)
	copy(out, s)

	//we are using a min heap of the max items in s
	//every time see an item smaller than the min item in s
	//we pop the min item from s and push our larger item
	h := IntHeap(out)
	heap.Init(&h)

	//O(n) iterations
	for i := len(out); i < len(s); i++ {
		if s[i] > h[0] {
			//O(log m)*2 = O(log m)
			heap.Pop(&h)
			heap.Push(&h, s[i])
		}
	}
	//O(m + n*log m))
	return out
}

func main() {
	a := make([]int, 100000)
	for i := range a {
		a[i] = rand.Intn(math.MaxInt32)
	}
	//
	ts := time.Now()
	if false {
		MaxElements(a, 20)
	} else {
		sort.Ints(a)
	}
	fmt.Println(time.Now().Sub(ts))
}
