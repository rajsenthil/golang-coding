package datastructures

import (
	"fmt"
	"log"
	"reflect"
)

type QueueNode struct {
	val  string
	next *QueueNode
}

type Queue struct {
	head *QueueNode
	tail *QueueNode
}

func (n *QueueNode) String() string {
	return fmt.Sprintf("Node: %s", n.val)
}

func (q *Queue) String() string {
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

func (q *Queue) queue(n *QueueNode) *QueueNode {
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

func (q *Queue) dequeue() QueueNode {
	curr := QueueNode{}
	if q == nil || &q.head == nil {
		return curr
	}
	curr = *q.head
	log.Printf("Current Head: %v", curr.String())
	prev := QueueNode{}
	for &curr != nil {
		if &curr != nil {
			prev = curr
			curr = *curr.next
		}
		log.Printf("Previous Node: %v", prev.String())
	}
	return QueueNode{}
}

func (q *Queue) dequeue2() *QueueNode {
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
