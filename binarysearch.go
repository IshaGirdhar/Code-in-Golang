package main

import (
	"fmt"
)

func BinarySearch(a []int, x int) bool {

	start := 0
	end := len(a)
	for start < end {
		mid := (start + end) / 2
		if a[mid] == x {
			return true
		} else if a[mid] < x {
			start = mid + 1
		} else {
			end = mid
		}

	}
	return false
}

func main() {
	a := []int{1, 2, 3, 4, 5, 7, 8, 10, 11}
	fmt.Println(BinarySearch(a, 4))

}
