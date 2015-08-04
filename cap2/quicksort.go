package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	input := os.Args[1:]
	numbers := make([]int, len(input))

	for i, n := range input {
		number, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s isn't a valid number\n", n)
		}
		numbers[i] = number
	}
	fmt.Printf("Input: %d\n", numbers)
	fmt.Printf("Output: %d\n", quicksort(numbers))
}

func quicksort(numbers []int) []int {
	if len(numbers) <= 1 {
		return numbers
	}

	// copy slice
	n := make([]int, len(numbers))
	copy(n, numbers)

	// select pivot
	indexPivot := len(n) / 2
	pivot := n[indexPivot]

	// remove pivot
	n = append(n[:indexPivot], n[indexPivot+1:]...)
	// fmt.Println(n)

	lower, bigger := partition(n, pivot)

	// return recursive
	return append(
		append(quicksort(lower), pivot),
		quicksort(bigger)...,
	)
}

func partition(numbers []int, pivot int) (lower []int, bigger []int) {
	for _, v := range numbers {
		if v <= pivot {
			lower = append(lower, v)
		} else {
			bigger = append(bigger, v)
		}
	}

	return lower, bigger
}
