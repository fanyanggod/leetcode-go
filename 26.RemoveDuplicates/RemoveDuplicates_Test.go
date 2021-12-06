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
			element{[]int{1, 1, 2}},
		},
		{
			element{[]int{1}},
		},
		{
			element{[]int{1, 3}},
		},
		{
			element{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
		},
	}
	for _, q := range qs {
		fmt.Println(removeDeuplicates(q.elements.nums), q.elements.nums)
	}
}
