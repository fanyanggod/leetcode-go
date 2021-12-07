package main

func searchInsert(nums []int, target int) int {
	l := len(nums)
	lp := 0
	rp := l - 1
	for lp <= rp {
		if nums[lp] < target {
			lp++
		}
		if nums[rp] >= target {
			rp--
		}
	}
	return lp
}
