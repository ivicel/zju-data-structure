package ch05_test

import (
	"fmt"
	"testing"

	"github.com/ivicel/zju_data_structure/ch03"
	"github.com/ivicel/zju_data_structure/ch05"
)

func TestHuffmanTree(t *testing.T) {
	arr := []interface{}{1, 2, 3, 4, 5}
	tree := ch05.CreateHuffmanTree(arr, func(n1, n2 *ch03.TreeNode) int {
		x := n1.Data.(int)
		y := n2.Data.(int)
		if x == y {
			return 0
		} else if y > x {
			return 1
		} else {
			return -1
		}
	})

	res := []interface{}{15, 6, 9, 3, 3, 4, 5, 1, 2}
	fmt.Printf("生成的 Huffman 树:")
	ch03.LevelTraversal(tree.Root)
	fmt.Printf("\n%v\n", res)
}
