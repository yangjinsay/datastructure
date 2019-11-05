package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	num1 := 0
	t1 := 1
	num2 := 0
	t2 := 1
	for {
		if l1.Next != nil {
			num1 = l1.Val * t1 + num1
			t1 = t1 * 10
			l1 = l1.Next
		} else {
			num1 = l1.Val * t1 + num1
			break
		}
	}


	for {
		if l2.Next != nil {
			num1 = l2.Val * t2 + num1
			t2 = t2 * 10
			l2 = l2.Next
		} else {
			num2 = l2.Val * t2 + num2
			break
		}
	}

	sum :=  num1 + num2
	result := &ListNode{}
	for {
		if sum/10 == 0 {
			return result
		} else {
			result.Next = &ListNode{}
			result.Val = sum%10
		}
	}
}

