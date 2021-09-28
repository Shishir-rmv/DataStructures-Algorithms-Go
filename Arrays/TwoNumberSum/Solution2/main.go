package main

import "fmt"

// O(n) time | O(n) space
func TwoNumberSum(array []int, target int) []int {
	m := make(map[int]bool, len(array))
	result := make([]int, 0)

	for _, v1 := range array {

		diff := target - v1
		if _, ok := m[diff]; ok {
			result = append(result, v1)
			result = append(result, diff)
		}

		if _, ok := m[v1]; !ok {
			m[v1] = true
		}
	}
	return result
}

func main() {
	array := []int{3, 5, -4, 8, 11, 1, -1, 6}
	targetSum := 10
	fmt.Println(TwoNumberSum(array, targetSum))
}
