package main

import "fmt"

type questions struct {
	elements element
}

type element struct {
	nums1 []int
	nums2 []int
}

func main() {
	qs := []questions{
		{
			element{[]int{1, 3}, []int{2}},
		},
		{},
		{
			element{[]int{1, 2}, []int{3, 4}},
		},
		{
			element{[]int{}, []int{1}},
		},
		{
			element{[]int{2}, []int{}},
		},
	}
	for _, q := range qs {
		fmt.Println(findMedianSortedArrays(q.elements.nums1, q.elements.nums2), q.elements.nums1, q.elements.nums2)
	}
}
