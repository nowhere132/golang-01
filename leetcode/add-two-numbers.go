package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curry := 0

	result := &ListNode{}
	p := result

	travel := func(p *ListNode) (int, *ListNode) {
		if p == nil {
			return 0, nil
		}
		return p.Val, p.Next
	}

	for {
		val1, nextL1 := travel(l1)
		val2, nextL2 := travel(l2)
		val := val1 + val2 + curry
		p.Val = val % 10
		l1 = nextL1
		l2 = nextL2
		//fmt.Println(l1)
		//fmt.Println(l2)

		curry = 0
		if val > 9 {
			curry = 1
		}
		//fmt.Println(val, p.Val, curry)

		if l1 == nil && l2 == nil {
			break
		}
		p.Next = &ListNode{}
		p = p.Next
	}

	if curry > 0 {
		p.Next = &ListNode{Val: curry, Next: nil}
	}

	return result
}

func (l *ListNode) Insert(x int) {
	l.Next = &ListNode{Val: x, Next: nil}
}
