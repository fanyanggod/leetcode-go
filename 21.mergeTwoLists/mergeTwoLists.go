package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ans *ListNode
	var current *ListNode
	for list1 != nil || list2 != nil {
		min := 0
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				min = list1.Val
				list1 = list1.Next
			} else {
				min = list2.Val
				list2 = list2.Next
			}
		} else if list1 == nil && list2 != nil {
			min = list2.Val
			list2 = list2.Next
		} else {
			min = list1.Val
			list1 = list1.Next
		}
		if ans == nil {
			ans = &ListNode{min, nil}
			current = ans
		} else {
			current.Next = &ListNode{min, nil}
			current = current.Next
		}
	}
	return ans
}
