package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node // 头结点
	len  int   // 长度
}

func main() {
	// 生成链表
	list := generate_list(20)
	fmt.Print("原始生成的列表: ")
	print_list(list)
	fmt.Printf("\n翻转后的列表: ")
	reverse(list, 21)
	print_list(list)
	fmt.Println()
}

// 翻转前 k 个数
func reverse(list *List, k int) {
	// 链表只有一个结点, 或只需要翻转一个结点
	if k == 1 || list.len == 1 {
		return
	}

	if k > list.len {
		k = list.len
	}

	optr := list.head // 指向未翻转的结点

	var tmp, // 指向已翻转的结点的下一临时结点
		nptr *Node // 已翻转的结点

	// 翻转前 k 个结点, 我们只需要移动 k - 1 次指针
	for i := 0; i < k-1; i++ {
		if i == 0 {
			tmp = list.head.next
		}

		nptr = tmp
		// 保证 tmp 总是指到下一个待翻转的结点, 顺序不能转
		tmp = nptr.next
		// 当前结点的 next 指向旧结点
		nptr.next = optr
		// 不要忘记移动旧结点
		optr = nptr
	}

	// 原头结点此时应该指向待翻转的结点
	list.head.next = tmp
	// 重新指向头结点
	list.head = optr

}

func generate_list(k int) *List {
	list := &List{}
	for i := k; i > 0; i-- {
		node := &Node{data: i}
		list.len++
		if i == k {
			list.head = node
			continue
		} else {
			node.next = list.head
			list.head = node
		}
	}

	return list
}

// 打印链表
func print_list(list *List) {
	ptr := list.head
	for {
		if ptr == nil {
			break
		}

		fmt.Printf("%d ", ptr.data)
		ptr = ptr.next
	}
}
