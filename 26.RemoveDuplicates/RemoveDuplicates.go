package main

func removeDeuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	lp := 1
	rp := 1
	for ; rp < l; rp++ {
		if nums[rp] != nums[rp-1] {
			nums[lp] = nums[rp]
			lp++
		}
	}
	return lp
}
