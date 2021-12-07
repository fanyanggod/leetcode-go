package main

func plusOne(digits []int) []int {
	l := len(digits)
	for i := l - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < l; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	digits = make([]int, l+1)
	digits[0] = 1
	return digits
}
