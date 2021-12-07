package main

import "fmt"

func maxSubArray(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	max := nums[0]

	current := nums[0]

	for i := 1; i < l; i++ {
		if nums[i] < 0 {
			if nums[i]+current > 0 {
				current += nums[i]
			} else {
				current = nums[i]
			}
		} else {
			if current < 0 {
				current = nums[i]
			} else {
				current += nums[i]
			}
		}
		if current > max {
			max = current
		}
		fmt.Println(current)
	}
	return max
}
