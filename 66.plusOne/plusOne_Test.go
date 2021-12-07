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
			element{[]int{1, 2, 2}},
		},
		{
			element{[]int{1}},
		},
		{
			element{[]int{4, 3}},
		},
		{
			element{[]int{9}},
		},
	}
	for _, q := range qs {
		fmt.Println(plusOne(q.elements.nums), q.elements.nums)
	}
}
