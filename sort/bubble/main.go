package main

import "fmt"

//for each neighboring pair of integers in the slice
//swap if left is larger than right
//stop when nothing was swapped
//complexity O(n2)
func bubbleSort(s []int) {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(s); i++ {
			if s[i-1] > s[i] {
				temp := s[i]
				s[i] = s[i-1]
				s[i-1] = temp
				swapped = true
			}
		}
	}
}

func main() {
	s := []int{1, 2, 4, 6, 2, 0, 7}
	bubbleSort(s)
	fmt.Println(s)
}
