package ch04

import (
	"github.com/ivicel/zju_data_structure/ch03"
)

type AVLTree struct {
	ch03.BinaryTree
	Equal func(*ch03.TreeNode, *ch03.TreeNode) int
}

// 插入新数据
func (t *AVLTree) Insert(value interface{}) {
	node := &ch03.TreeNode{
		Depth: 1,
		Data:  value,
	}

	t.Root = t.insertUnderNode(t.Root, node)
}

func (t *AVLTree) insertUnderNode(parent, newNode *ch03.TreeNode) *ch03.TreeNode {
	// 如果是一棵空树时, 直接返回
	// 如果该结点是已经是叶子结点, 则返回新生成的结点
	if parent == nil {
		return newNode
	}

	// 将新结点插入到左子树还是右子树
	if t.Equal(parent, newNode) < 0 {
		parent.Right = t.insertUnderNode(parent.Right, newNode)
	} else {
		parent.Left = t.insertUnderNode(parent.Left, newNode)
	}

	// 对比当前结点的左右子树的高度, 判断是否需要重新调整平衡
	if abs(getHeight(parent.Left)-getHeight(parent.Right)) > 1 {
		parent = t.makeBalance(parent)
	}

	t.calculateHeight(parent)
	return parent
}

func (t *AVLTree) makeBalance(node *ch03.TreeNode) *ch03.TreeNode {
	var node1, node2 *ch03.TreeNode
	if getHeight(node.Left) > getHeight(node.Right) &&
		getHeight(node.Left.Left) >= getHeight(node.Left.Right) {
		// ll 失衡
		node1 = node.Left
		node.Left = node1.Right
		node1.Right = node
		t.calculateHeight(node)
	} else if getHeight(node.Right) > getHeight(node.Left) &&
		getHeight(node.Right.Right) >= getHeight(node.Right.Left) {
		// rr 失衡
		node1 = node.Right
		node.Right = node1.Left
		node1.Left = node
		t.calculateHeight(node)
	} else if getHeight(node.Left) > getHeight(node.Right) &&
		getHeight(node.Left.Right) > getHeight(node.Left.Left) {
		// lr 失衡
		// 先做左旋转, 再做右旋转
		node1 = node.Left.Right
		node2 = node.Left
		node.Left = node1.Right
		node2.Right = node1.Left
		node1.Right = node
		node1.Left = node2

		t.calculateHeight(node)
		t.calculateHeight(node2)
	} else if getHeight(node.Right) > getHeight(node.Left) &&
		getHeight(node.Right.Left) > getHeight(node.Right.Right) {
		// rl 失衡
		// 先做右旋转, 再做左旋转
		node1 = node.Right.Left
		node2 = node.Right
		node.Right = node1.Left
		node2.Left = node1.Right
		node1.Left = node
		node1.Right = node2

		t.calculateHeight(node)
		t.calculateHeight(node2)
	}

	return node1
}

func (t *AVLTree) calculateHeight(node *ch03.TreeNode) {
	left := getHeight(node.Left)
	right := getHeight(node.Right)
	node.Depth = max(left, right) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func getHeight(node *ch03.TreeNode) int {
	if node == nil {
		return 0
	} else {
		return node.Depth
	}
}
