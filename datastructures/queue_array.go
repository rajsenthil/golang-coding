package datastructures

import (
  "errors"
  "fmt"
  "log"
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
  if n != nil {
    return fmt.Sprintf("Node val: %s", n.val)
  } else {
    return ""
  }
}

func (q *QueueArray) String() string {
  s := ""
  if q.array[0] == nil {
    return s
  }
  for i := 0; i <= q.tailPos; i++ {
    s = s + fmt.Sprint(q.array[i].String()) + " ==> "
  }
  return s
}

func (q *QueueArray) queueArray(n *QueueArrayNode) (*QueueArray, error) {
  if q.tailPos >= QUEUE_SIZE-1 {
    return nil, errors.New("Queue reached its capacity")
  }
  q.tailPos = q.tailPos + 1
  q.array[q.tailPos] = n
  log.Printf("Tail position after adding: %d", q.tailPos)
  return q, nil
}

func (q *QueueArray) dequeueArray() (*QueueArrayNode, error) {
  if len(q.array) == 0 || q.tailPos < 0 {
    return nil, errors.New("Queue is empty")
  }
  log.Printf("Length of array: %d", len(q.array))
  head := q.array[0]
  for i := 0; i < len(q.array)-1; i++ {
    q.array[i] = q.array[i+1]
  }
  q.tailPos = q.tailPos - 1
  log.Printf("Queue after dequeueing head: %v is : %v", head, q.String())
  return head, nil
}
