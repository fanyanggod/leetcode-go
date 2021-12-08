package main

import "fmt"

type questions struct {
	l1 ListNode
}

type element struct {
	nums []int
}

func main() {

	qs := []questions{
		{
			ListNode{2, &ListNode{4, &ListNode{3, nil}}},
		},
		{
			ListNode{1, &ListNode{2, &ListNode{4, nil}}},
		},
		{
			ListNode{0, nil},
		},
	}
	for _, q := range qs {
		fmt.Println(swapPairs(&q.l1), q.l1)
	}
}
