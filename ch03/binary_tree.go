package ch03

import (
	"fmt"

	"github.com/ivicel/zju_data_structure/ch02"
)

// 二叉树
type BinaryTree struct {
	Root *TreeNode // 头结点
	Len  int       // 树的长度
}

// 二叉树结点
type TreeNode struct {
	Depth int         // 该节点的深度
	Data  interface{} // 节点数据
	Left  *TreeNode   // 左子结点
	Right *TreeNode   // 右子结点
}

func (n TreeNode) String() string {
	return fmt.Sprintf("%v", n.Data)
}

// 前序-递归实现
func RecusivPreOrderTraversal(root *TreeNode) {
	if root != nil {
		fmt.Printf("%v ", root.Data)
		RecusivPreOrderTraversal(root.Left)
		RecusivPreOrderTraversal(root.Right)
	}
}

// 中序-递归实现
func RecusivInOrderTraversal(root *TreeNode) {
	if root != nil {
		RecusivInOrderTraversal(root.Left)
		fmt.Printf("%v ", root.Data)
		RecusivInOrderTraversal(root.Right)
	}
}

// 后序-递归实现
func RecusivPostOrderTraversal(root *TreeNode) {
	if root != nil {
		RecusivPostOrderTraversal(root.Left)
		RecusivPostOrderTraversal(root.Right)
		fmt.Printf("%v ", root.Data)
	}
}

// 前序-非递归栈实现
func StackPreOrderTraversal(root *TreeNode) {
	stack := &ch02.Stack{}
	for root != nil || !stack.IsEmpty() {
		for root != nil {
			fmt.Printf("%v ", root.Data)
			stack.Push(&ch02.StackNode{Data: root})
			root = root.Left
		}

		if !stack.IsEmpty() {
			node := stack.Pop()
			root = node.Data.(*TreeNode).Right
		}
	}
}

// 中序-非递归栈实现
func StackInOrderTraversal(n *TreeNode) {
	stack := &ch02.Stack{}
	for n != nil || !stack.IsEmpty() {
		// 找到最底层的左子树叶子结点
		for n != nil {
			stack.Push(&ch02.StackNode{Data: n})
			n = n.Left
		}

		// 现在栈里最顶结点就是树里最左的叶子结点的父结点
		if !stack.IsEmpty() {
			n1 := stack.Pop().Data.(*TreeNode)
			fmt.Printf("%v ", n1.Data)
			n = n1.Right
		}
	}
}

// // 单个栈实现后序遍历
func StackPostOrderTraversal(n *TreeNode) {
	s1 := &ch02.Stack{}
	var pre *TreeNode
	for n != nil || !s1.IsEmpty() {
		for n != nil {
			s1.Push(&ch02.StackNode{Data: n})
			n = n.Left
		}

		if !s1.IsEmpty() {
			// 经过上面循环, 以某个结点为子树的话, 栈中要么是父结点, 要么是左结点
			// 如果有右子结点, 那说明这个一个父结点, 根据后序左右根规则, 需要将父结点重新入栈
			// 然后转到右子树, 入栈右子树, 然后这个右结点会被入栈
			// 假设某个右叶子结点, 在此时出栈, 那么其符合条件 1, 此时记录下这个结点
			// 然后下一次出栈的就是它的父结点, 此时对比父结点的右结点 == 上次我们记录的结点
			n1 := s1.Pop().Data.(*TreeNode)
			if n1.Right == nil || pre == n1.Right {
				fmt.Printf("%v ", n1.Data)
				pre = n1
				// 当左结点没有了右子结点, 那说明这个结点是左叶子结点, 则应该访问左叶子结点
				// 此时应该把结点置空, 这样下次循环时便跳过
				n = nil
			} else {
				// 重新入栈父结点
				s1.Push(&ch02.StackNode{Data: n1})
				n = n1.Right
			}
		}
	}
}

func DoubleStackPostOrderTraversal(n *TreeNode) {
	s1 := &ch02.Stack{}
	s2 := &ch02.Stack{}
	for n != nil || !s1.IsEmpty() {
		for n != nil {
			s1.Push(&ch02.StackNode{Data: n})
			s2.Push(&ch02.StackNode{Data: n})
			n = n.Right
		}

		if !s1.IsEmpty() {
			n1 := s1.Pop().Data.(*TreeNode)
			n = n1.Left
		}
	}

	for !s2.IsEmpty() {
		n1 := s2.Pop().Data.(*TreeNode)
		fmt.Printf("%s ", n1.Data)
	}
}

// 层级遍历
func LevelTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	queue := ch02.Queue{}
	queue.Offer(&ch02.QueueNode{Data: root})
	for {
		n := queue.Poll()
		if n == nil {
			break
		}

		n1 := n.Data.(*TreeNode)
		fmt.Printf("%v ", n1.Data)

		if n1.Left != nil {
			queue.Offer(&ch02.QueueNode{Data: n1.Left})
		}

		if n1.Right != nil {
			queue.Offer(&ch02.QueueNode{Data: n1.Right})
		}
	}
}
