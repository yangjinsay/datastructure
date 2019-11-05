/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
    Val int
    Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 {
		return head
	}
	getAllLN(head)
	le := len(allNode)
	if le == 1 {
		return nil
	}
	allNode[le - n -1] = allNode[le - n + 1]
	return head
}
var allNode  []*ListNode

func getAllLN(head *ListNode){
	allNode = append(allNode, head)
	if head == nil {
		return
	}
	getAllLN(head.Next)
}
