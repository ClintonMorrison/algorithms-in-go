package main

import "fmt"

func main() {
	nums := []int{1, 2, 2, 5, 8, 10}
	result := binarySearch(nums, 8)
	fmt.Println(result)
}

func binarySearch(nums []int, search int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := (lo + hi) / 2
		if (nums[mid] < search) {
			lo = mid + 1
		} else if (nums[mid] > search) {
			hi = mid - 1
		} else {
			return mid
		}
	}

	return -1
}
