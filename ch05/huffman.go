package ch05

import (
	"github.com/ivicel/zju_data_structure/ch03"
)

type HuffmanTree struct {
	ch03.BinaryTree
}

func CreateHuffmanTree(arr []interface{}, cmp EqualFunc) *HuffmanTree {
	queue := CreateQueue(arr, cmp)
	tree := &HuffmanTree{}
	for queue.Size > 0 {
		n1, _ := queue.Pop()
		n2, _ := queue.Pop()
		parent := ch03.TreeNode{Data: n1.Data.(int) + n2.Data.(int)}
		if cmp(&n1, &n2) > 0 {
			parent.Left = &n1
			parent.Right = &n2
		} else {
			parent.Left = &n2
			parent.Right = &n1
		}

		if queue.Size == 0 {
			tree.Root = &parent
			break
		}
		_ = queue.Push(parent)
	}

	return tree
}
