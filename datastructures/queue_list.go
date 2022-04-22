package datastructures

import (
  "errors"
  "fmt"
  "log"
  "reflect"
)

type QueueListNode struct {
  val  string
  next *QueueListNode
}

type QueueList struct {
  head *QueueListNode
  tail *QueueListNode
}

func (n *QueueListNode) String() string {
  if n == nil {
    return fmt.Sprintf("Node: NONE")
  }
  return fmt.Sprintf("Node: %s", n.val)
}

func (q *QueueList) String() string {
  log.Printf("Printing Queue")
  node := q.head
  log.Printf("Head: %v", node.String())
  s := ""
  for node != nil {
    s += fmt.Sprintf("Node %s => ", node.val)
    node = node.next
  }
  s += " nil"
  return s
}

func (q *QueueList) queueList(n *QueueListNode) *QueueListNode {
  if q.head == nil || q.head.val == "" {
    q.head = n
    q.tail = n
    return n
  }
  curr := q.head
  tail := q.head
  for curr == nil || curr.val != "" {
    tail = curr
    curr = curr.next
    if curr == nil {
      break
    }
  }
  tail.next = n
  q.tail = tail
  log.Printf(q.String())
  return q.head
}

func (q *QueueList) dequeueList() (*QueueListNode, error) {
  head := q.head
  if head == nil {
    return head, errors.New("Empty Queue")
  }
  q.head = q.head.next
  return head, nil
}

func (q *QueueList) dequeueListFromTail() *QueueListNode {
  curr := q.head.next
  prev := q.head
  log.Printf("curr type: %v", reflect.ValueOf(curr).Kind())
  log.Printf("prev type: %v", reflect.ValueOf(prev).Kind())
  log.Printf("Starting Prev: %v", prev)
  log.Printf("Starting curr: %v", curr)
  temp := curr
  for temp != nil {
    temp = curr.next
    if temp == nil {
      break
    }
    prev = curr
    curr = curr.next
    log.Printf("Running Prev: %v", prev)
    log.Printf("Running curr: %v", curr)
  }
  log.Printf("Prev Node: %v", prev.String())
  prev.next = nil
  q.tail = prev
  log.Printf("Tail node: %v", q.tail.String())
  log.Printf("Queue: %v", q.String())
  return curr
}
