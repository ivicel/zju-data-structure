package ch04_test

import (
	"testing"

	"github.com/ivicel/zju_data_structure/ch03"
	"github.com/ivicel/zju_data_structure/ch04"
)

// 比较两个结点的大小
// 相等返回 0, 小于返回 -1, 大于返回 1
func Equal(n1, n2 *ch03.TreeNode) int {
	x := n1.Data.(int)
	y := n2.Data.(int)
	if x == y {
		return 0
	} else if x > y {
		return 1
	} else {
		return -1
	}
}

func TestAvlTree(t *testing.T) {
	data := [...]int{16, 3, 7, 11, 9, 26, 18, 14, 15}
	tree := &ch04.AVLTree{Equal: Equal}
	for _, d := range data {
		tree.Insert(d)
	}

	ch03.RecusivPreOrderTraversal(tree.Root)
}
