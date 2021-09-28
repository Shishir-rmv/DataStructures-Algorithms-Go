package main

import (
	"fmt"
	"sort"
)

// O(nlog(n)) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)

	left, right := 0, len(array)-1
	for left < right {
		sum := array[left] + array[right]
		if sum == target {
			return []int{array[left], array[right]}
		} else if sum < target {
			left += 1
		} else {
			right -= 1
		}
	}

	return []int{}
}

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	fmt.Println(TwoNumberSum(array, targetSum))
}
