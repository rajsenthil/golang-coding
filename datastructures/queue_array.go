package datastructures

import (
	"errors"
	"fmt"
)

type QueueArrayNode struct {
	val string
}

const (
	QUEUE_SIZE = 10
)

type QueueArray struct {
	tailPos int
	array   [QUEUE_SIZE]*QueueArrayNode
}

func (n *QueueArrayNode) String() string {
	return fmt.Sprintf("Node val: %s", n.val)
}

func (q *QueueArray) String() string {
	s := ""
	if q.array[0] == nil {
		return s
	}
	for i := 0; i < q.tailPos; i++ {
		s = s + fmt.Sprintf(q.array[i].String()) + " ==> "
	}
	return s
}

func (q *QueueArray) queueArray(n *QueueArrayNode) (*QueueArray, error) {
	if q.tailPos >= QUEUE_SIZE-1 {
		return nil, errors.New("Queue reached its capacity")
	}
	q.array[q.tailPos] = n
	q.tailPos = q.tailPos + 1
	return q, nil
}

func (q *QueueArray) dequeueArray() (*QueueArrayNode, error) {
	if len(q.array) == 0 {
		return nil, errors.New("Queue is empty")
	}

	head := q.array[0]
	for i := 1; i < len(q.array)-1; i++ {
		q.array[i] = q.array[i+1]
		q.tailPos = q.tailPos - 1
	}
	return head, nil
}
