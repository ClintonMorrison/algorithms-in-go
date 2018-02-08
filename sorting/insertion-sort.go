package main

import "fmt"

func main() {
	nums := []int{1, 6, 4, 2, 9, 5, 8, 4, 1, 0}
	insertionSort(nums)
	fmt.Println(nums)
}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i += 1 {
		insertIntoSorted(nums, i-1, nums[i])
	}
}

func insertIntoSorted(nums []int, sortedUpToIndex int, numToInsert int) {
	i := sortedUpToIndex

	// Shift numbers in sorted portion right until a number less than numToInsert is found
	for i >= 0 && nums[i] > numToInsert {
		nums[i+1] = nums[i]
		i -= 1
	}

	nums[i+1] = numToInsert
}
