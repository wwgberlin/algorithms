package main

import (
	"fmt"
)

func main() {
	s := []int{5, 2, 3, 5, 9, 1, 2, 2}
	fmt.Println(mergeSort(s))
	fmt.Println(nonRecursiveMergeSort(s))
}

func mergeSort(unsorted []int) []int {
	out := make([]int, len(unsorted))
	return recursiveMergeSort(unsorted, out)
}

func recursiveMergeSort(unsorted []int, out []int) []int {
	if len(unsorted) < 2 {
		return unsorted
	}
	l := len(unsorted) / 2
	merge(
		recursiveMergeSort(unsorted[:l], out),
		recursiveMergeSort(unsorted[l:], out),
		out)
	copy(unsorted, out)
	return unsorted
}

func merge(lhs []int, rhs []int, out []int) {
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
	return
}

func nonRecursiveMergeSort(unsorted []int) []int {
	out := make([]int, len(unsorted))
	for k := 1; k <= len(unsorted); k *= 2 {
		for i := 0; i < len(unsorted); i += k {
			lhs := unsorted[i: i+(k/2)]
			rhs := unsorted[i+(k/2):i+k]
			merge(lhs, rhs, out)
			copy(unsorted[i:i+k], out)
		}
	}
	return unsorted
}
