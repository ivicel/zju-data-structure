package main

import (
	"errors"
	"fmt"
)

// 树结点
type Node struct {
	val int
}

// 大小比较
func (node *Node) Equal(other *Node) int {
	if node.val == other.val {
		return 0
	} else if node.val > other.val {
		return 1
	} else {
		return -1
	}
}

// 最大堆
type PriorityQueue struct {
	size    int    // 当前队列大小
	maxSize int    // 队列的最大容量
	val     []Node // 存储数据结点
}

// 插入新的结点
func (q *PriorityQueue) Push(val int) error {
	if q.size == q.maxSize {
		return errors.New("队列已满, 插入失败")
	}

	// 将结点放到最后位置
	newNode := Node{val}
	q.size++
	q.val[q.size] = newNode
	// 上滤
	pos := q.size
	for q.val[0] = newNode; newNode.Equal(&q.val[pos/2]) > 0; pos = pos / 2 {
		// 大于时交换值
		q.val[pos], q.val[pos/2] = q.val[pos/2], q.val[pos]
	}

	return nil
}

// 删除最大值
func (q *PriorityQueue) Pop() (Node, error) {
	if q.size == 0 {
		return Node{}, errors.New("不能删除空树")
	}

	// 保存根结点
	node := q.val[1]
	q.val[1] = q.val[q.size]
	q.size--

	q.percolateDown(1)

	return node, nil
}

// 下滤
func (q *PriorityQueue) percolateDown(root int) {
	q.val[0] = q.val[root]
	// 如果左孩子都不存在的话, 那就是叶子结点了, 因为这是一棵完全二叉树
	for root*2 <= q.size {
		pos, rightPos := root*2, root*2+1

		// 如果右子树存在并且右 > 左, 那 pos = 右, 否则 pos = 左
		if rightPos <= q.size && q.val[pos].Equal(&q.val[rightPos]) < 0 {
			pos = rightPos
		}

		// 比较孩子结点和父结点的大小, pos 可能是左也可能是右
		// 不大于则说明最大结点就是父结点
		if q.val[pos].Equal(&q.val[0]) > 0 {
			q.val[root] = q.val[pos]
			root = pos
		} else {
			// 最大值就是父结点时, 不用交换
			break
		}
	}
	q.val[root] = q.val[0]
}

// 打印树内容
func (q *PriorityQueue) String() string {
	return fmt.Sprintf("%v", q.val[1:q.size+1])
}

// 初始化一个空堆
func InitQueue(maxSize int) *PriorityQueue {
	queue := &PriorityQueue{size: 0, maxSize: maxSize}
	queue.val = make([]Node, maxSize)
	return queue
}

// 根据组生成一个新堆
func CreateQueue(arr []int) *PriorityQueue {
	// 按默认顺序创建一个新堆
	size := len(arr)
	maxSize := size * 2
	queue := &PriorityQueue{
		size:    size,
		maxSize: maxSize,
		val:     make([]Node, 1, maxSize),
	}

	nodes := make([]Node, size)
	for i, val := range arr {
		nodes[i] = Node{val}
	}
	queue.val = append(queue.val, nodes...)

	// 调整排序
	for i := size / 2; i > 0; i-- {
		queue.percolateDown(i)
	}

	return queue
}

func main() {
	arr := [...]int{79, 66, 16, 83, 30, 19, 68, 55, 91, 72, 49, 9}
	queue := InitQueue(20)
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
	for queue.size > 0 {
		if node, err := queue.Pop(); err != nil {
			fmt.Printf("无法删除结点: [%s]\n", err.Error())
			break
		} else {
			fmt.Printf("删除结点: %v\n", node)
		}
	}

	if queue.size > 0 {
		fmt.Printf("当前堆中还有剩余结点: %v\n", queue)
	}

	fmt.Printf("删除完成\n\n-----------------------\n\n")

	fmt.Println("时间复杂度 O(N) 生成新的堆")
	queue = CreateQueue(arr[:])
	fmt.Println(queue)
}
