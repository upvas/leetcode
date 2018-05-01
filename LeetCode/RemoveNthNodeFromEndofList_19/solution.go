package RemoveNthNodeFromEndofList_19


type ListNode struct {
	Val int
	Next *ListNode
}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var cnt int
	for el := head; el != nil; el = el.Next {
		cnt++
	}

	if n == cnt { // remove the first element
		return head.Next
	}

	var cur int
	var el *ListNode
	for el = head; el != nil; el = el.Next {
		if cur == cnt - n - 1 {
			break
		}
		cur++
	}
	el.Next = el.Next.Next

	return head
}