package main

import "fmt"

type questions struct {
	elements element
}

type element struct {
	nums    []int
	tartget int
}

func main() {
	qs := []questions{
		{
			element{[]int{1, 1, 2}, 0},
		},
		{
			element{[]int{1}, 2},
		},
		{
			element{[]int{1, 3}, 2},
		},
		{
			element{[]int{1, 3, 5, 6}, 5},
		},
	}
	for _, q := range qs {
		fmt.Println(searchInsert(q.elements.nums, q.elements.tartget), q.elements.nums, q.elements.tartget)
	}
}
