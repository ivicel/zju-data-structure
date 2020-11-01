package ch03_test

import (
	"fmt"
	"testing"

	"github.com/ivicel/zju_data_structure/ch03"
)

// 生成树
//			     A
//			  /    \
//		   B          C
//       /  \       /   \
//      D    F     G      I
//          /       \
//         E         H
func generateTree() *ch03.BinaryTree {
	return &ch03.BinaryTree{
		Len: 9,
		Root: &ch03.TreeNode{
			Data: "A",
			Left: &ch03.TreeNode{
				Data: "B",
				Left: &ch03.TreeNode{
					Data: "D",
				},
				Right: &ch03.TreeNode{
					Data: "F",
					Left: &ch03.TreeNode{
						Data: "E",
					},
				},
			},
			Right: &ch03.TreeNode{
				Data: "C",
				Left: &ch03.TreeNode{
					Data: "G",
					Right: &ch03.TreeNode{
						Data: "H",
					},
				},
				Right: &ch03.TreeNode{
					Data: "I",
				},
			},
		},
	}
}

func TestBinaryTree(t *testing.T) {
	bt := generateTree()
	fmt.Printf("使用递归的前序遍历: ")
	ch03.RecusivPreOrderTraversal(bt.Root)
	fmt.Println()

	fmt.Printf("使用递归的中序遍历: ")
	ch03.RecusivInOrderTraversal(bt.Root)
	fmt.Println()

	fmt.Printf("使用递归的后序遍历: ")
	ch03.RecusivPostOrderTraversal(bt.Root)
	fmt.Println()

	fmt.Printf("使用栈的前序遍历: ")
	ch03.StackPreOrderTraversal(bt.Root)
	fmt.Println()

	fmt.Printf("使用栈的中序遍历: ")
	ch03.StackInOrderTraversal(bt.Root)
	fmt.Println()

	fmt.Printf("使用栈的后序遍历: ")
	ch03.StackPostOrderTraversal(bt.Root)
	fmt.Println()

	fmt.Printf("使用双栈的后序遍历: ")
	ch03.DoubleStackPostOrderTraversal(bt.Root)
	fmt.Println()

	fmt.Printf("使用队列的层级遍历: ")
	ch03.LevelTraversal(bt.Root)
	fmt.Println()
}
