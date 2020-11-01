/**
 * 队列
 */

package ch02

// 双端队列
type Queue struct {
	Head *QueueNode // 队列头
	Tail *QueueNode // 队列尾
}

// 双端队列结点
type QueueNode struct {
	Data interface{} // 储存的数据
	Next *QueueNode  // 队列下一结点
}

// Offer 添加到队尾
func (q *Queue) Offer(newNode *QueueNode) {
	if q.Tail == nil {
		q.Head = newNode
		q.Tail = newNode
	} else {
		q.Tail.Next = newNode
		q.Tail = newNode
	}
}

// 从队列中取出
func (q *Queue) Poll() *QueueNode {
	if q.Head == nil {
		return nil
	}

	node := q.Head
	q.Head = q.Head.Next
	if q.Head == nil {
		q.Tail = nil
	}
	return node
}
