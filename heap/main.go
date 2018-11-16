package main

import (
	"fmt"
	"math"
)

//the parent of 0 is 0
//and somehow that doesn't bother the algorithms :/
func parent(i int) int {
	return (i - 1) / 2
}
func left(i int) int {
	return i*2 + 1
}
func right(i int) int {
	return i*2 + 2
}

/*
From Introduction to Algorithms (Corman):
MAX-HEAPIFY(A, i)
1 	l = LEFT(i)
2 	r = RIGHT(i)
3 	if l <= A.heap-size and A[l] > A[i]
4 		largest = l
5 	else largest = i
6 	if r <= A.heap-size and A[r] > A[largest]
7 		largest = r
8 	if largest != i
9 		exchange A[i] with A[largest]
10 		MAX-HEAPIFY(A, largest)
*/
func maxHeapify(a []int, i int) {
	l := left(i)
	r := right(i)
	var largest int
	if l < len(a) && a[l] > a[i] {
		largest = l
	} else {
		largest = i
	}
	if r < len(a) && a[r] > a[largest] {
		largest = r
	}
	if largest != i {
		swap(a, i, largest)
		maxHeapify(a, largest)
	}
}

func swap(a []int, i, j int) {
	if i == j {
		return
	}
	//example: a[i] = 3, a[j] = 5
	a[i] = a[i] + a[j] //a[i] is now 8
	a[j] = a[i] - a[j] //a[j] is now 8-5=3
	a[i] = a[i] - a[j] //a[j] is now 8-3=5
}

/*
BUILD-MAX-HEAP(A)
From Introduction to Algorithms (Corman):
1 	A.heap-size := A.length
2 	for i := [A.length/2] downto 1
3 	MAX-HEAPIFY(A, i)
*/

func buildMaxHeap(a []int) {
	//n-1 iterations
	//but wait, what is going on with the complexity here?
	//This is an O(n) operation (iteration)
	//times O(log n) =
	//O(n*log n)
	//Why is this O(n)?
	//It is actually because of the property of n that
	//we are passing to maxHeapify, it is not n
	//it is h <= n/2^(h+1)
	for i := len(a) / 2; i >= 0; i-- {
		maxHeapify(a, i)
	}
}

/*
From Introduction to Algorithms (Corman):
HEAPSORT(A)
1 BUILD-MAX-HEAP(A)
2 for i := A.length downto 2
3 	exchange A[1] with A[i]
4 	A.heap-size := A.heap-size - 1
5 	MAX-HEAPIFY(A, 1)/
*/
func heapSort(a []int) {
	buildMaxHeap(a)
	for i := len(a) - 1; i >= 0; i-- {
		swap(a, 0, i)
		a = a[:len(a)-1]
		maxHeapify(a, 0)
	}
}

/*
From Introduction to Algorithms (Corman):
HEAP-INCREASE-KEY(A, i, key)
1 if key < A[i]
2 	error “new key is smaller than current key”
3 A[i] = key
4 while i>1 and A[PARENT(i)] < A[i]
5 	exchange A[i] with A(PARENT[i]) /
6 	i = PARENT(i) /
*/
func heapIncreaseKey(a []int, i int, key int) {
	if key < a[i] {
		//we will get to this
		panic("new key is smaller than current key")
	}
	a[i] = key
	for ; i > 0 && a[parent(i)] < a[i]; i = parent(i) {
		swap(a, i, parent(i))
	}
}

/*
From Introduction to Algorithms (Corman):
MAX-HEAP-INSERT(A, key)
1 A.heap-size = A.heap-size + 1
2 A[A.heap-size] = -infinity
3 HEAP-INCREASE-KEY(A, A.heap-size, key)
*/
func maxHeapInsert(a *[]int, key int) {
	n := append(*a, math.MinInt32)
	*a = n
	heapIncreaseKey(*a, len(*a)-1, key)
}

func main() {
	a := []int{4, 1, 3, 2, 16, 9, 10, 14, 8, 7}
	buildMaxHeap(a)
	fmt.Println(a)
	maxHeapInsert(&a, 11)
	maxHeapInsert(&a, -1)
	fmt.Println(a)
	heapSort(a)
	fmt.Println(a)
}
