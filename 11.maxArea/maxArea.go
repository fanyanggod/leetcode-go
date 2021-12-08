package main

func maxArea(height []int) int {
	lowPointer := 0
	heightPointer := len(height) - 1
	ans := 0
	for lowPointer <= heightPointer {
		if height[lowPointer] <= height[heightPointer] {
			if height[lowPointer]*(heightPointer-lowPointer) >= ans {
				ans = height[lowPointer] * (heightPointer - lowPointer)
			}
			lowPointer++
		} else {
			if height[heightPointer]*(heightPointer-lowPointer) >= ans {
				ans = height[heightPointer] * (heightPointer - lowPointer)
			}
			heightPointer--
		}
	}
	return ans
}
