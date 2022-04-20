package datastructures

import (
	"log"
	"testing"
)

func TestAddNodeToQueue(t *testing.T) {
	node1 := QueueArrayNode{val: "1"}
	node2 := QueueArrayNode{val: "2"}
	node3 := QueueArrayNode{val: "3"}
	node4 := QueueArrayNode{val: "4"}
	node5 := QueueArrayNode{val: "5"}
	queue := QueueArray{}
	queue.queueArray(&node1)
	queue.queueArray(&node2)
	queue.queueArray(&node3)
	queue.queueArray(&node4)
	queue.queueArray(&node5)
	log.Printf("Queue: %s", queue.String())
}

func TestAddNodeToQueueMoreThan10(t *testing.T) {
	node0 := QueueArrayNode{val: "0"}
	node1 := QueueArrayNode{val: "1"}
	node2 := QueueArrayNode{val: "2"}
	node3 := QueueArrayNode{val: "3"}
	node4 := QueueArrayNode{val: "4"}
	node5 := QueueArrayNode{val: "5"}
	node6 := QueueArrayNode{val: "6"}
	node7 := QueueArrayNode{val: "7"}
	node8 := QueueArrayNode{val: "8"}
	node9 := QueueArrayNode{val: "9"}
	node10 := QueueArrayNode{val: "10"}
	queue := QueueArray{}
	queue.queueArray(&node0)
	queue.queueArray(&node1)
	queue.queueArray(&node2)
	queue.queueArray(&node3)
	queue.queueArray(&node4)
	queue.queueArray(&node5)
	queue.queueArray(&node6)
	queue.queueArray(&node7)
	queue.queueArray(&node8)
	_, err := queue.queueArray(&node9)
	if err == nil {
		log.Printf("Logic not working. Error raised before filling the capacity: %v", err)
		return
	}
	_, err = queue.queueArray(&node10)
	if err != nil {
		log.Printf("Error raised as expected: %v", err)
	} else {
		log.Print("Logic not working")
		return
	}
	log.Printf("Queue: %s", queue.String())
}
