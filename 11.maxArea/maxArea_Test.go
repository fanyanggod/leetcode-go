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
			element{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
		},
		{
			element{[]int{1, 1}},
		},
		{
			element{[]int{4, 3, 2, 1, 4}},
		},
		{
			element{[]int{1, 2, 1}},
		},
	}
	for _, q := range qs {
		fmt.Println(maxArea(q.elements.nums), q.elements.nums)
	}
}
