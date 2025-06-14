package main

// 206. Reverse Linked List
type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	current := head
	for current != nil {
		temp := current.Next
		current.Next = prev
		prev = current
		current = temp
	}
	return prev
}

// 21. Merge Two Sorted Lists
func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
	}
	if l1 != nil {
		current.Next = l1
	} else {
		current.Next = l2
	}
	return dummy.Next
}

// 141. Linked List Cycle
func HasCycle(head *ListNode) bool {
	// if head == nil {
	// 	return false
	// }
	// fast := head
	// slow := head
	// for fast != nil && fast.Next != nil && slow != nil {
	// 	fast = fast.Next.Next
	// 	slow = slow.Next
	// 	if fast == slow {
	// 		return true
	// 	}
	// }
	// return false

	visited := make(map[*ListNode]bool)
	node := head
	for node != nil {
		if _, exists := visited[node]; exists {
			return true
		}
		visited[node] = true
		node = node.Next
	}
	return false
}

// 2. Add Two Numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		d1, d2 := 0, 0
		if l1 != nil {
			d1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			d2 = l2.Val
			l2 = l2.Next
		}
		digit := d1 + d2 + carry
		remainder := digit % 10
		carry = digit / 10

		current.Next = &ListNode{remainder, nil}
		current = current.Next

	}
	if carry != 0 {
		current.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
