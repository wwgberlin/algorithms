package main

import "fmt"

func main() {
	s := []int{5, 2, 3, 5, 9, 1, 2, 2}
	fmt.Println(mergeSort(s))
	fmt.Println(nonRecursiveMergeSort(s))
}

//Recursively merge sort given unsorted slice.
//out base check is if len(unsorted) is 0 or 1, then we have nothing to sort
//and we simply return unsorted.
//If we have more than 1 item we divide the unsorted slice in the middle into
//two slices, we call mergeSort on each, and then merge the two slices using
// the given function merge that takes two sorted slices and merges them.
func mergeSort(unsorted []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}
	l := len(unsorted) / 2
	return merge(mergeSort(unsorted[:l]), mergeSort(unsorted[l:]))
}

func merge(lhs []int, rhs []int) []int {
	out := make([]int, len(lhs)+len(rhs))
	ilhs, irhs := 0, 0

	for i := 0; i < len(lhs)+len(rhs); i++ {
		if ilhs >= len(lhs) {
			out[i] = rhs[irhs]
			irhs++
		} else if irhs >= len(rhs) {
			out[i] = lhs[ilhs]
			ilhs++
		} else if lhs[ilhs] < rhs[irhs] {
			out[i] = lhs[ilhs]
			ilhs++
		} else {
			out[i] = rhs[irhs]
			irhs++
		}
	}
	return out
}

func nonRecursiveMergeSort(unsorted []int) []int {
	for k := 1; k <= len(unsorted); k *= 2 {
		for i := 0; i < len(unsorted); i += k {
			lhs := unsorted[i: i+(k/2)]
			rhs := unsorted[i+(k/2):i+k]
			copy(unsorted[i:i+k], merge(lhs, rhs))
		}
	}
	return unsorted
}
