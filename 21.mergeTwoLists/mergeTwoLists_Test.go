package main

import "fmt"

type questions struct {
	l1 ListNode
	l2 ListNode
}

type element struct {
	nums []int
}

func main() {

	qs := []questions{
		{
			ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			ListNode{5, &ListNode{6, &ListNode{9, nil}}},
		},
		{
			ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			ListNode{1, &ListNode{3, &ListNode{4, nil}}},
		},
		{
			ListNode{},
			ListNode{0, nil},
		},
	}
	for _, q := range qs {
		fmt.Println(mergeTwoLists(&q.l1, &q.l2), q.l1, q.l2)
	}
}
