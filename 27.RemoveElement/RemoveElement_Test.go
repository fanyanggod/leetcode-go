package main

import "fmt"

type questions struct {
	elements element
}

type element struct {
	nums []int
	val  int
}

func main() {
	qs := []questions{
		{
			element{[]int{1, 1, 2}, 1},
		},
		{
			element{[]int{1}, 1},
		},
		{
			element{[]int{1, 3}, 1},
		},
		{
			element{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 3},
		},
	}
	for _, q := range qs {
		fmt.Println(removeElement(q.elements.nums, q.elements.val), q.elements.nums)
	}
}
