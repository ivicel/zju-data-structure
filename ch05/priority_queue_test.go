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
	arr := []interface{}{79, 66, 43, 83, 30, 87, 38, 55, 91, 72, 49, 9}
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

	res := []interface{}{91, 87, 83, 79, 72, 43, 38, 55, 66, 30, 49, 9}
	for i, j := range res {
		if j != queue.Data[i+1].Data {
			t.Fatalf("新建堆结果错误, 应为: %v, 实际为: %v\n", j, queue.Data[i+1])
		}
	}

	fmt.Printf("插入完成: %v\n", queue)

	fmt.Printf("\n-----------------------\n\n")

	res = []interface{}{91, 87, 83, 79, 72, 66, 55, 49, 43, 38, 30, 9}
	i := 0
	fmt.Println("删除结点演示")
	for queue.Size > 0 {
		if node, err := queue.Pop(); err != nil {
			fmt.Printf("无法删除结点: [%s]\n", err.Error())
			break
		} else if res[i] != node.Data {
			t.Fatalf("删除结点错误, 应为: %v, 实际为: %v\n", res[i], node)
		}
		i++
	}
	fmt.Printf("%v\n", res)

	if queue.Size > 0 {
		fmt.Printf("当前堆中还有剩余结点: %v\n", queue)
	}

	fmt.Printf("删除完成\n\n-----------------------\n\n")

	res = []interface{}{91, 83, 87, 79, 72, 43, 38, 55, 66, 30, 49, 9}
	queue = ch05.CreateQueue(arr, Equal)
	fmt.Printf("时间复杂度 O(N) 生成新的堆: %+v\n", queue)
	for i, j := range res {
		if j != queue.Data[i+1].Data {
			t.Fatalf("新建堆结果错误, 应为: %v, 实际为: %v\n", j, queue.Data[i+1])
		}
	}
	// fmt.Println(queue.Data)
}
