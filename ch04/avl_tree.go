package main

import "fmt"

// 树的节点
type Node struct {
	depth int   // 该节点的深度
	val   int   // 节点数据
	left  *Node // 左子结点
	right *Node // 右子结点
}

// 比较两个结点的大小
// 相等返回 0, 小于返回 -1, 大于返回 1
func (node *Node) Equal(other *Node) int {
	if node.val == other.val {
		return 0
	} else if node.val > other.val {
		return 1
	} else {
		return -1
	}
}

type AVLTree struct {
	root *Node // 根结点
}

// 插入新数据
func (t *AVLTree) Insert(value int) {
	node := &Node{
		depth: 1,
		val:   value,
	}

	t.root = t.insertUnderNode(t.root, node)
}

func (t *AVLTree) insertUnderNode(parent *Node, newNode *Node) *Node {
	// 如果是一棵空树时, 直接返回
	// 如果该结点是已经是叶子结点, 则返回新生成的结点
	if parent == nil {
		return newNode
	}

	// 将新结点插入到左子树还是右子树
	if parent.Equal(newNode) < 0 {
		parent.right = t.insertUnderNode(parent.right, newNode)
	} else {
		parent.left = t.insertUnderNode(parent.left, newNode)
	}

	// 对比当前结点的左右子树的高度, 判断是否需要重新调整平衡
	if abs(getHeight(parent.left)-getHeight(parent.right)) > 1 {
		parent = t.makeBalance(parent)
	}

	t.calculateHeight(parent)
	return parent
}

func (t *AVLTree) makeBalance(node *Node) *Node {
	var node1, node2 *Node
	if getHeight(node.left) > getHeight(node.right) &&
		getHeight(node.left.left) >= getHeight(node.left.right) {
		// ll 失衡
		node1 = node.left
		node.left = node1.right
		node1.right = node
		t.calculateHeight(node)
	} else if getHeight(node.right) > getHeight(node.left) &&
		getHeight(node.right.right) >= getHeight(node.right.left) {
		// rr 失衡
		node1 = node.right
		node.right = node1.left
		node1.left = node
		t.calculateHeight(node)
	} else if getHeight(node.left) > getHeight(node.right) &&
		getHeight(node.left.right) > getHeight(node.left.left) {
		// lr 失衡
		// 先做左旋转, 再做右旋转
		node1 = node.left.right
		node2 = node.left
		node.left = node1.right
		node2.right = node1.left
		node1.right = node
		node1.left = node2

		t.calculateHeight(node)
		t.calculateHeight(node2)
	} else if getHeight(node.right) > getHeight(node.left) &&
		getHeight(node.right.left) > getHeight(node.right.right) {
		// rl 失衡
		// 先做右旋转, 再做左旋转
		node1 = node.right.left
		node2 = node.right
		node.right = node1.left
		node2.left = node1.right
		node1.left = node
		node1.right = node2

		t.calculateHeight(node)
		t.calculateHeight(node2)
	}

	return node1
}

func (t *AVLTree) calculateHeight(node *Node) {
	left := getHeight(node.left)
	right := getHeight(node.right)
	node.depth = max(left, right) + 1
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

func getHeight(node *Node) int {
	if node == nil {
		return 0
	} else {
		return node.depth
	}
}

func PreOrderTraversal(node *Node) {
	if node != nil {
		fmt.Printf("%d ", node.val)
		PreOrderTraversal(node.left)
		PreOrderTraversal(node.right)
	}
}

func main() {
	data := [...]int{16, 3, 7, 11, 9, 26, 18, 14, 15}
	tree := &AVLTree{}
	for _, d := range data {
		tree.Insert(d)
	}

	PreOrderTraversal(tree.root)
	fmt.Println()
}
