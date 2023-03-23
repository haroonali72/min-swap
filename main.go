package main

import "fmt"

func main() {
	fmt.Println(minSwaps([]int32{4, 3, 1, 2}))
}

func minSwaps(nums []int32) int32 {
	var (
		swaps  int32
		i      int32
		length = int32(len(nums))
	)

	for i = 0; i < length; i++ {
		minIndex := minNumIndex(nums[i:])
		if nums[i] > nums[minIndex+i] {
			temp := nums[i]
			nums[i] = nums[minIndex+i]
			nums[minIndex+i] = temp

			swaps = swaps + 1
		}
	}

	return swaps
}

func minNumIndex(nums []int32) int32 {
	var (
		min   = nums[0]
		index int32
	)

	for i := 0; i < len(nums); i++ {
		if min > nums[i] {
			min = nums[i]
			index = int32(i)
		}
	}

	return index
}
