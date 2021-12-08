package main

import "fmt"

type questions struct {
	elements element
}

type element struct {
	nums []int
}

func main() {
	qs := []questions{
		{
			element{[]int{-1, 0, 1, 2, -1, -4}},
		},
		{
			element{[]int{}},
		},
		{
			element{[]int{0}},
		},
		{
			element{[]int{-2, 0, 0, 2, 2}},
		},
	}
	for _, q := range qs {
		fmt.Println(threeSum(q.elements.nums), q.elements.nums)
	}
}
