package ch02

import (
	"fmt"
)

// 单向链表
type List struct {
	Head *ListNode // 头结点
	Len  int   // 长度
}

// 链表节结
type ListNode struct {
	Data interface{}
	Next *ListNode
}

// 翻转前 k 个数
func Reverse(list *List, k int) {
	// 链表只有一个结点, 或只需要翻转一个结点
	if k == 1 || list.Len == 1 {
		return
	}

	if k > list.Len {
		k = list.Len
	}

	optr := list.Head // 指向未翻转的结点

	var tmp, // 指向已翻转的结点的下一临时结点
		nptr *ListNode // 已翻转的结点

	// 翻转前 k 个结点, 我们只需要移动 k - 1 次指针
	for i := 0; i < k-1; i++ {
		if i == 0 {
			tmp = list.Head.Next
		}

		nptr = tmp
		// 保证 tmp 总是指到下一个待翻转的结点, 顺序不能转
		tmp = nptr.Next
		// 当前结点的 next 指向旧结点
		nptr.Next = optr
		// 不要忘记移动旧结点
		optr = nptr
	}

	// 原头结点此时应该指向待翻转的结点
	list.Head.Next = tmp
	// 重新指向头结点
	list.Head = optr
}

// 打印链表
func PrintList(list *List) {
	ptr := list.Head
	for {
		if ptr == nil {
			break
		}

		fmt.Printf("%d ", ptr.Data)
		ptr = ptr.Next
	}
}
