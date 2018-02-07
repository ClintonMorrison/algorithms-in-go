package main

import "fmt"

func main() {
  nums := []int{1, 6, 4, 2, 9, 5, 8, 4, 1, 0}
  bubblesort(nums)
  fmt.Println(nums)
}

func bubblesort(array []int) {
  for {
    swapped := false
    for i := 0; i < len(array) - 1; i += 1 {
      if (array[i] > array[i+1]) {
        array[i], array[i+1] = array[i+1], array[i]
        swapped = true
      }
    }

    if !swapped {
      return;
    }
  }
}

