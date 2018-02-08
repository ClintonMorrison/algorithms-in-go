package main

import "fmt"

func main() {
	nums := []int{1, 6, 4, 2, 9, 5, 8, 4, 1, 0}
	insertionSort(nums)
	fmt.Println(nums)

}

func insertionSort(nums []int) {
	for i := 1; i < len(nums); i += 1 {
		numToInsert := nums[i]

		j := i - 1
		for ; j >= 0 && nums[j] >= numToInsert; j -= 1 {
			nums[j + 1] = nums[j]
		}

		nums[j+1] = numToInsert
	}
}
