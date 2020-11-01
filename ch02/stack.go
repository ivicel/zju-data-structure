/**
 * 链表实现的栈
 */

package ch02

// 栈
type Stack struct {
	Head *StackNode
}

type StackNode struct {
	Head *StackNode
	Data interface{}
	Next *StackNode
}

// 压入栈
func (s *Stack) Push(n *StackNode) {
	if n != nil {
		if s.Head == nil {
			s.Head = n
		} else {
			n.Next = s.Head
			s.Head = n
		}
	}
}

// 弹出栈
func (s *Stack) Pop() *StackNode {
	if s.Head == nil {
		return nil
	} else {
		t := s.Head
		s.Head = s.Head.Next
		return t
	}
}

// 判空
func (s *Stack) IsEmpty() bool {
	return s.Head == nil
}
