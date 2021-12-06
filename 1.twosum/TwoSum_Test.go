package main

import "fmt"

type questions struct {
	elements element
	anser    []int
}

type element struct {
	nums   []int
	target int
}

func main() {
	qs := []questions{
		{
			element{[]int{2, 7, 11, 15}, 9},
			[]int{0, 1},
		},
		{
			element{[]int{3, 2, 4}, 6},
			[]int{1, 2},
		},
		{
			element{[]int{3, 3}, 6},
			[]int{0, 1},
		},
	}
	for _, q := range qs {
		fmt.Println(twoSum(q.elements.nums, q.elements.target))
	}
}
