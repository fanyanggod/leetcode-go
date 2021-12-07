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
			element{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
		},
		{
			element{[]int{1}},
		},
		{
			element{[]int{1, -1}},
		},
		{
			element{[]int{5, 4, -1, 7, 8}},
		},
		{
			element{[]int{-2, -1}},
		},
		{
			element{[]int{-2, -1}},
		},
		{
			element{[]int{-2, 0}},
		},
	}
	for _, q := range qs {
		fmt.Println(maxSubArray(q.elements.nums), q.elements.nums)
	}
}
