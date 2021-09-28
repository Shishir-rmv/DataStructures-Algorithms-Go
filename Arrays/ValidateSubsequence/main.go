package main

import "fmt"

func validateSubsequence(array []int, sequence []int) bool {
	i := 0
	j := 0

	for i < len(array) && j < len(sequence) {
		if sequence[j] == array[i] {
			j++
		}
		i++
	}

	return j == len(sequence)
}

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	sequence := []int{1, 6, -1, 10}
	fmt.Println(validateSubsequence(array, sequence))
}
