package ch02_test

import (
	"fmt"
	"testing"

	"github.com/ivicel/zju_data_structure/ch02"
)

func TestReverseList(t *testing.T) {
	// 生成链表
	list := generateList(20)
	fmt.Print("原始生成的列表: ")
	ch02.PrintList(list)
	fmt.Printf("\n翻转后的列表: ")
	ch02.Reverse(list, 21)
	ch02.PrintList(list)
	fmt.Println()
}

func generateList(k int) *ch02.List {
	list := &ch02.List{}
	for i := k; i > 0; i-- {
		node := &ch02.ListNode{Data: i}
		list.Len++
		if i == k {
			list.Head = node
			continue
		} else {
			node.Next = list.Head
			list.Head = node
		}
	}

	return list
}
