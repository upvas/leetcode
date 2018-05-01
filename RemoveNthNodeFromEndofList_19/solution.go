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

// One pass with slow and fast markers
func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	dummy := &ListNode{}
	dummy.Next = head

	slow := dummy
	fast := dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}