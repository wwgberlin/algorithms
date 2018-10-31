package main

import "fmt"

func main() {
	s := []int{5, 2, 3, 5, 9, 1, -1, 0, -4, 2}
	quicksort(s)
	fmt.Println("result: ",s)
}

func quicksort(unsorted []int) {
	if len(unsorted) > 1 {
		p := partition(unsorted)
		quicksort(unsorted[:p])
		quicksort(unsorted[p+1:])
	}
}

func swap(unsorted []int, i, j int) {
	temp := unsorted[i]
	unsorted[i] = unsorted[j]
	unsorted[j] = temp
}

func partition(unsorted []int) int {
	fmt.Println("in: ", unsorted)
	pivot := unsorted[len(unsorted)-1]
	i := 0
	for j := 0; j < len(unsorted)-1; j++ {
		if unsorted[j] < pivot {
			if i != j {
				swap(unsorted, i, j)
			}
			i++
		}
	}
	swap(unsorted, i, len(unsorted)-1)
	fmt.Println("out:", unsorted)
	return i
}
