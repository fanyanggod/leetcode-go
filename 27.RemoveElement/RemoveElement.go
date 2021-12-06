package main

func removeElement(nums []int, val int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	lp := 0
	rp := l - 1

	for rp >= lp {
		if nums[rp] == val {
			rp--
		} else {
			if nums[lp] == val {
				nums[lp] = nums[rp]
				nums[rp] = val
			}
			lp++
		}
	}
	return lp
}
