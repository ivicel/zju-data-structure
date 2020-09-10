package main

import "fmt"

// 二叉树
type BinaryTree struct {
	root *Node // 头结点
	len  int   // 树的长度
}

type StackNode struct {
	data *Node
	next *StackNode
}

// 栈
type Stack struct {
	head *StackNode
}

// 压入栈
func (s *Stack) push(node *StackNode) {
	if node != nil {
		if s.head == nil {
			s.head = node
		} else {
			node.next = s.head
			s.head = node
		}
	}
}

// 弹出栈
func (s *Stack) pop() *StackNode {
	if s.head == nil {
		return nil
	} else {
		t := s.head
		s.head = s.head.next
		return t
	}
}

// 判空
func (s *Stack) isEmpty() bool {
	return s.head == nil
}

type Node struct {
	data  string // 数据
	left  *Node  // 左子结点
	right *Node  // 右子结点
}

// 前序-递归实现
func recusivPreOrderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%s ", root.data)
		recusivPreOrderTraversal(root.left)
		recusivPreOrderTraversal(root.right)
	}
}

// 中序-递归实现
func recusivInOrderTraversal(root *Node) {
	if root != nil {
		recusivInOrderTraversal(root.left)
		fmt.Printf("%s ", root.data)
		recusivInOrderTraversal(root.right)
	}
}

// 后序-递归实现
func recusivPostOrderTraversal(root *Node) {
	if root != nil {
		recusivPostOrderTraversal(root.left)
		recusivPostOrderTraversal(root.right)
		fmt.Printf("%s ", root.data)
	}
}

// 前序-非递归栈实现
func stackPreOrderTraversal(node *Node) {
	stack := &Stack{}
	for node != nil || !stack.isEmpty() {
		for node != nil {
			fmt.Printf("%s ", node.data)
			stack.push(&StackNode{data: node})
			node = node.left
		}

		if !stack.isEmpty() {
			stackNode := stack.pop()
			node = stackNode.data.right
		}
	}
}

// 中序-非递归栈实现
func stackInOrderTraversal(node *Node) {
	stack := &Stack{}
	for node != nil || !stack.isEmpty() {
		// 找到最底层的左子树叶子结点
		for node != nil {
			stack.push(&StackNode{data: node})
			node = node.left
		}

		// 现在栈里最顶结点就是树里最左的叶子结点的父结点
		if !stack.isEmpty() {
			stackNode := stack.pop()
			fmt.Printf("%s ", stackNode.data.data)
			node = stackNode.data.right
		}
	}
}

func stackPostOrderTraversal(node *Node) {
	s1 := &Stack{}
	var pre *Node
	for node != nil || !s1.isEmpty() {
		for node != nil {
			s1.push(&StackNode{data: node})
			node = node.left
		}

		if !s1.isEmpty() {
			// 经过上面循环, 以某个结点为子树的话, 栈中要么是父结点, 要么是左结点
			// 如果有右子结点, 那说明这个一个父结点, 根据后序左右根规则, 需要将父结点重新入栈
			// 然后转到右子树, 入栈右子树, 然后这个右结点会被入栈
			// 假设某个右叶子结点, 在此时出栈, 那么其符合条件 1, 此时记录下这个结点
			// 然后下一次出栈的就是它的父结点, 此时对比父结点的右结点 == 上次我们记录的结点
			stackNode := s1.pop()
			if stackNode.data.right == nil || pre == stackNode.data.right {
				fmt.Printf("%s ", stackNode.data.data)
				pre = stackNode.data
				// 当左结点没有了右子结点, 那说明这个结点是左叶子结点, 则应该访问左叶子结点
				// 此时应该把结点置空, 这样下次循环时便跳过
				node = nil
			} else {
				// 重新入栈父结点
				s1.push(stackNode)
				node = stackNode.data.right
			}
		}
	}
}

func doubleStackPostOrderTraversal(node *Node) {
	s1 := &Stack{}
	s2 := &Stack{}
	for node != nil || !s1.isEmpty() {
		for node != nil {
			s1.push(&StackNode{data: node})
			s2.push(&StackNode{data: node})
			node = node.right
		}

		if !s1.isEmpty() {
			stackNode := s1.pop()
			node = stackNode.data.left
		}
	}

	for !s2.isEmpty() {
		fmt.Printf("%s ", s2.pop().data.data)
	}
}

// 链表队列
type Queue struct {
	head *QueueNode // 队列头
	tail *QueueNode // 队列尾
}

// 链表结点
type QueueNode struct {
	data *Node
	next *QueueNode
}

// 添加到队首
func (q *Queue) offer(node *QueueNode) {
	if q.tail == nil {
		q.head = node
		q.tail = node
	} else {
		q.tail.next = node
		q.tail = node
	}
}

// 从队列中取出
func (q *Queue) poll() *QueueNode {
	if q.head == nil {
		return nil
	}

	node := q.head
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return node
}

// 层级遍历
func levelTraversal(root *Node) {
	if root == nil {
		return
	}

	queue := Queue{}
	queue.offer(&QueueNode{data: root})
	for {
		node := queue.poll()
		if node == nil {
			break
		}
		fmt.Printf("%s ", node.data.data)
		if node.data.left != nil {
			queue.offer(&QueueNode{data: node.data.left})
		}

		if node.data.right != nil {
			queue.offer(&QueueNode{data: node.data.right})
		}
	}
}

func main() {
	bt := generate_tree()
	fmt.Printf("使用递归的前序遍历: ")
	recusivPreOrderTraversal(bt.root)
	fmt.Println()

	fmt.Printf("使用递归的中序遍历: ")
	recusivInOrderTraversal(bt.root)
	fmt.Println()

	fmt.Printf("使用递归的后序遍历: ")
	recusivPostOrderTraversal(bt.root)
	fmt.Println()

	fmt.Printf("使用栈的前序遍历: ")
	stackPreOrderTraversal(bt.root)
	fmt.Println()

	fmt.Printf("使用栈的中序遍历: ")
	stackInOrderTraversal(bt.root)
	fmt.Println()

	fmt.Printf("使用栈的后序遍历: ")
	stackPostOrderTraversal(bt.root)
	fmt.Println()

	fmt.Printf("使用双栈的后序遍历: ")
	doubleStackPostOrderTraversal(bt.root)
	fmt.Println()

	fmt.Printf("使用队列的层级遍历: ")
	levelTraversal(bt.root)
	fmt.Println()
}

// 生成树
//			     A
//			  /    \
//		   B          C
//       /  \       /   \
//      D    F     G      I
//          /       \
//         E         H
func generate_tree() *BinaryTree {
	return &BinaryTree{
		len: 9,
		root: &Node{
			data: "A",
			left: &Node{
				data: "B",
				left: &Node{
					data: "D",
				},
				right: &Node{
					data: "F",
					left: &Node{
						data: "E",
					},
				},
			},
			right: &Node{
				data: "C",
				left: &Node{
					data: "G",
					right: &Node{
						data: "H",
					},
				},
				right: &Node{
					data: "I",
				},
			},
		},
	}
}
