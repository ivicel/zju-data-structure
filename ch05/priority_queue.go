package ch05

import (
	"errors"
	"fmt"

	"github.com/ivicel/zju_data_structure/ch03"
)

type EqualFunc func(*ch03.TreeNode, *ch03.TreeNode) int

// 最大堆
type PriorityQueue struct {
	Size    int             // 当前队列大小
	MaxSize int             // 队列的最大容量
	Data    []ch03.TreeNode // 存储数据结点
	Equal   EqualFunc
}

// 插入新的结点
func (q *PriorityQueue) Push(val interface{}) error {
	if q.Size == q.MaxSize {
		return errors.New("队列已满, 插入失败")
	}

	// 将结点放到最后位置
	newNode := ch03.TreeNode{Data: val}
	q.Size++
	q.Data[q.Size] = newNode
	// 上滤
	pos := q.Size
	for q.Data[0] = newNode; q.Equal(&newNode, &q.Data[pos/2]) > 0; pos = pos / 2 {
		// 大于时交换值
		q.Data[pos], q.Data[pos/2] = q.Data[pos/2], q.Data[pos]
	}

	return nil
}

// 删除最大值
func (q *PriorityQueue) Pop() (ch03.TreeNode, error) {
	if q.Size == 0 {
		return ch03.TreeNode{}, errors.New("不能删除空树")
	}

	// 保存根结点
	node := q.Data[1]
	q.Data[1] = q.Data[q.Size]
	q.Size--

	q.percolateDown(1)

	return node, nil
}

// 下滤
func (q *PriorityQueue) percolateDown(root int) {
	q.Data[0] = q.Data[root]
	// 如果左孩子都不存在的话, 那就是叶子结点了, 因为这是一棵完全二叉树
	for root*2 <= q.Size {
		pos, rightPos := root*2, root*2+1

		// 如果右子树存在并且右 > 左, 那 pos = 右, 否则 pos = 左
		if rightPos <= q.Size && q.Equal(&q.Data[pos], &q.Data[rightPos]) < 0 {
			pos = rightPos
		}

		// 比较孩子结点和父结点的大小, pos 可能是左也可能是右
		// 不大于则说明最大结点就是父结点
		if q.Equal(&q.Data[pos], &q.Data[0]) > 0 {
			q.Data[root] = q.Data[pos]
			root = pos
		} else {
			// 最大值就是父结点时, 不用交换
			break
		}
	}
	q.Data[root] = q.Data[0]
}

// 打印树内容
func (q *PriorityQueue) String() string {
	return fmt.Sprintf("%v", q.Data[1:q.Size+1])
}

// 初始化一个空堆
func InitQueue(maxSize int, cmp EqualFunc) *PriorityQueue {
	queue := &PriorityQueue{Size: 0, MaxSize: maxSize, Equal: cmp}
	queue.Data = make([]ch03.TreeNode, maxSize)
	return queue
}

// 根据传入的切片数据生成一个新堆
func CreateQueue(arr []interface{}, cmp EqualFunc) *PriorityQueue {
	// 按默认顺序创建一个新堆
	size := len(arr)
	maxSize := size * 2
	queue := InitQueue(maxSize, cmp)

	nodes := make([]ch03.TreeNode, size, maxSize)
	for i, val := range arr {
		nodes[i] = ch03.TreeNode{Data: val}
	}
	queue.Data = nodes

	// 调整排序
	for i := size / 2; i > 0; i-- {
		queue.percolateDown(i)
	}

	return queue
}
