package main

import "fmt"

func main() {
	s := []int{5, 2, 3, 5, 9, 1, -1, 0, -4, 2}
	quicksort(s)
	fmt.Println("result: ", s)

	//s = []int{5, 2, 3, 5, 9, 1, -1, 0, -4, 2}
	//quickSortIterative(s)
	//fmt.Println("result: ", s)
}

/*
From wikipedia (Quicksort):
The recursive quicksort pseudo code

algorithm quicksort(A, lo, hi) is
    if lo < hi then
        p := partition(A, lo, hi)
        quicksort(A, lo, p - 1 )
        quicksort(A, p + 1, hi)

In Go we can slice our slices (and our arrays)
So we don't have to pass into the function the upper and lower bound.
Instead we need to check the length of our slice and stop when
there's nothing to sort (length < 2)
To slice given slice a from 0 to l (not including l): a[:l]
To slice given slice a from l to the end of a: a[l:]
 */
func quicksort(unsorted []int) {
	if len(unsorted) < 2 {
		return
	}
	p:=partition(unsorted)
	quicksort(unsorted[:p])
	quicksort(unsorted[p+1:])
}

func swap(arr []int, i, j int) {
	temp := arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}

func partition(arr []int) int {
	//fmt.Println("in: ", arr)
	l := 0
	h := len(arr) - 1
	pivot := arr[h]
	//fmt.Println("pivot:", pivot)
	i := l - 1

	for j := l; j <= h-1; j++ {
		if arr[j] < pivot {
			i++
			swap(arr, i, j)
		}
	}
	swap(arr, i+1, h)
	//fmt.Println("out:", arr)
	return i + 1
}

type pair struct {
	left  int
	right int
}

func quickSortIterative(arr []int) {
	l := 0
	r := len(arr)

	var top int
	stack := make([]pair, len(arr))
	stack[top] = pair{left: l, right: r}

	for top >= 0 {
		l, r = stack[top].left, stack[top].right
		top--

		p := partition(arr[l: r])

		//p is the index of the pivot in the inner slice, add l to get
		//its index in the outer slice
		if p+l > l {
			top++
			stack[top] = pair{left: l, right: p + l}
		}

		if p+l+1 < r {
			top++
			stack[top] = pair{left: p + l + 1, right: r}
		}
	}
}
