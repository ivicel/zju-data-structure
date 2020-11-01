package ch05_test

import (
	"fmt"
	"testing"

	"github.com/ivicel/zju_data_structure/ch03"
	ch05 "github.com/ivicel/zju_data_structure/ch05"
)

// 大小比较
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

func TestPriorityQueue(t *testing.T) {
	arr := []interface{}{79, 66, 16, 83, 30, 19, 68, 55, 91, 72, 49, 9}
	queue := ch05.InitQueue(20, Equal)
	// 可以使用逐个插入新数据来生成一个新的堆, 时间复杂度是 O(NlogN)
	// 这是因为的插入 N 次, 每次调用堆复杂度是 logN
	fmt.Println("插入结点演示:")
	for _, val := range arr {
		if err := queue.Push(val); err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Printf("插入新结点 [%d]\n", val)
	}

	fmt.Printf("插入完成: %v\n", queue)
	fmt.Printf("\n-----------------------\n\n")

	fmt.Println("删除结点演示")
	for queue.Size > 0 {
		if node, err := queue.Pop(); err != nil {
			fmt.Printf("无法删除结点: [%s]\n", err.Error())
			break
		} else {
			fmt.Printf("删除结点: %v\n", node)
		}
	}

	if queue.Size > 0 {
		fmt.Printf("当前堆中还有剩余结点: %v\n", queue)
	}

	fmt.Printf("删除完成\n\n-----------------------\n\n")

	fmt.Println("时间复杂度 O(N) 生成新的堆")
	queue = ch05.CreateQueue(arr, Equal)
	fmt.Println(queue.Data)
}
