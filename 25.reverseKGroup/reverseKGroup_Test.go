package main

import "fmt"

type questions struct {
	l1 ListNode
	k  int
}

type element struct {
	nums []int
}

func main() {

	qs := []questions{
		{
			ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 3,
		},
		//{
		//	ListNode{1, &ListNode{2, &ListNode{4, nil}}}, 2,
		//},
		//{
		//	ListNode{0, nil}, 1,
		//},
	}
	for _, q := range qs {
		fmt.Println(reverseKGroup(&q.l1, q.k))
	}
}
