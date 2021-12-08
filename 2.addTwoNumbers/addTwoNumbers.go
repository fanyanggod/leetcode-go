package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var ans *ListNode
	var current *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		sum, carry = sum%10, sum/10
		if ans == nil {
			ans = &ListNode{sum, nil}
			current = ans
		} else {
			current.Next = &ListNode{sum, nil}
			current = current.Next
		}
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return ans
}
