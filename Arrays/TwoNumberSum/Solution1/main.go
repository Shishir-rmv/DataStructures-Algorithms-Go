package main

import "fmt"

// O(n^2) time | O(1) space
func TwoNumberSum(array []int, target int) []int {
	result := make([]int, 0)

	for i, v1 := range array {
		for k, v2 := range array {
			if i != k {
				sum := v1 + v2
				if sum == target {
					result = append(result, v1)
					result = append(result, v2)
					return result
				}
			}
		}
	}

	return result
}

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	fmt.Println(TwoNumberSum(array, targetSum))
}
