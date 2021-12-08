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
	}
	for _, q := range qs {
		fmt.Println(addTwoNumbers(&q.l1, &q.l2), q.l1, q.l2)
	}
}
