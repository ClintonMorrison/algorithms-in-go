package main

import "fmt"

func main() {
	nums := []int{1, 6, 4, 2, 9, 5, 8, 4, 1, 0}
	fmt.Println(mergeSort(nums))
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2

	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, len(left) + len(right))

	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result[l+r] =left[l]
			l += 1
		} else {
			result[l+r] = right[r]
			r += 1
		}
	}

	for l < len(left) {
		result[l+r] =left[l]
		l += 1
	}

	for r < len(right) {
		result[l+r] = right[r]
		r += 1
	}

	return result
}
