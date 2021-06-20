package main

 type ListNode struct {
 	Val int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	if head.Val == val {
		head = head.Next
		return head
	}
	tmp := head

	for tmp != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
			return head
		}
		tmp = tmp.Next
	}
	return head
}
