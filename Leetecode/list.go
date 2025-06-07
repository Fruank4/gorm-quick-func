package Leetecode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	p, q := headA, headB
	for p != q {
		if p == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	m, n := 0, 0
	p, q := headA, headB
	for p != nil {
		p = p.Next
		m++
	}
	for q != nil {
		q = q.Next
		n++
	}

	p, q = headA, headB
	if m > n {
		for i := 0; i < m-n; i++ {
			p = p.Next
		}
	} else {
		for i := 0; i < n-m; i++ {
			q = q.Next
		}
	}

	for p != nil && q != nil {
		if p == q {
			return p
		}
		p = p.Next
		q = q.Next
	}

	return nil
}

// 1, 2, 3, 4, 5

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	p, q := head, head.Next
	p.Next = nil
	newHead := reverseList(q) // 要用q，而不是
	q.Next = p
	return newHead
}

/*
*

	1 2 2 3 2 2 1
*/
func isPalindrome(head *ListNode) bool {
	mid := getMid(head)
	tail := reverseList(mid)
	p, q := head, tail
	p.Next = nil
	for p != nil && q != nil {
		if p.Val != q.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	return true
}

func getMid(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	// 1 2 3 4 4 5 6 7
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	return slow
}

func hasCycle(head *ListNode) bool {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}

func detectCycle(head *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			break
		}
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	helper := dummy
	for helper != slow {
		helper = helper.Next
		slow = slow.Next
	}
	return slow
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	p, q := l1, l2
	var next int
	for p != nil && q != nil {
		value := p.Val + q.Val + next
		cur := value % 10
		if value >= 10 {
			next = 1
		} else {
			next = 0
		}
		p.Val, q.Val = cur, cur
		if p.Next == nil && q.Next == nil && next != 0 {
			p.Next = &ListNode{
				Val:  next,
				Next: nil,
			}
			return l1
		}
		p = p.Next
		q = q.Next
	}

	if p != nil {
		for p != nil {
			value := p.Val + next
			cur := value % 10
			if value >= 10 {
				next = 1
			} else {
				next = 0
			}
			p.Val = cur
			if p.Next == nil && next != 0 {
				p.Next = &ListNode{
					Val:  next,
					Next: nil,
				}
				return l1
			}
			p = p.Next
		}
		return l1
	}

	for q != nil {
		for q != nil {
			value := q.Val + next
			cur := value % 10
			if value >= 10 {
				next = 1
			} else {
				next = 0
			}
			q.Val = cur
			if q.Next == nil && next != 0 {
				q.Next = &ListNode{
					Val:  next,
					Next: nil,
				}
				return l2
			}
			q = q.Next
		}
		return l2
	}

	return l1
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	pre, slow, fast := dummy, head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		pre = pre.Next
		slow = slow.Next
		fast = fast.Next
	}
	pre.Next = slow.Next
	return dummy.Next
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	mp := make(map[*Node]*Node)
	dummy := &Node{
		Val:    0,
		Next:   nil,
		Random: nil,
	}

	p := dummy
	q := head
	for q != nil {
		p.Next = &Node{
			Val:    q.Val,
			Random: q.Random,
			Next:   nil,
		}
		p = p.Next
		mp[q] = p
		q = q.Next
	}

	k := dummy.Next
	for k != nil {
		k.Random = mp[k.Random]
		k = k.Next
	}
	return dummy.Next
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	post := dummy
	for i := 0; i < k && post != nil; i++ {
		post = post.Next
	}
	if post == nil {
		return head
	}
	next := post.Next
	post.Next = nil

	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	dummy.Next.Next = reverseKGroup(next, k)
	return pre
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	if len(lists) == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}

	return mergeTwoLists(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	slow, fast := dummy, dummy
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	mid := slow.Next
	slow.Next = nil
	leftList := sortList(head) //
	rightList := sortList(mid) //
	return mergeTwoLists(leftList, rightList)
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	var cur = dummy
	p, q := list1, list2
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
	}
	if p != nil {
		cur.Next = p
	}

	if q != nil {
		cur.Next = q
	}
	return dummy.Next
}
