package algorithms

import (
	"errors"
	"log"
	"sync"
)

type Stack struct {
	lock  sync.Mutex
	items []byte
}

func NewStack() *Stack {
	return &Stack{items: make([]byte, 0)}
}

func (stack *Stack) push(item byte) {
	stack.lock.Lock()
	defer stack.lock.Unlock()
	stack.items = append(stack.items, item)
}

func (stack *Stack) pop() (byte, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	items_len := len(stack.items)
	if items_len == 0 {
		return 0, errors.New("Empty Stack")
	}
	result := stack.items[items_len-1]
	stack.items = stack.items[:items_len-1]
	return result, nil
}

func (stack *Stack) peek(item byte) (byte, error) {
	stack.lock.Lock()
	defer stack.lock.Unlock()

	items_len := len(stack.items)
	if items_len == 0 {
		return 0, errors.New("Empty Stack")
	}
	result := stack.items[items_len-1]
	return result, nil
}

func (stack *Stack) print() {
	items_len := len(stack.items)
	for idx := items_len - 1; idx >= 0; idx-- {
		log.Printf("items[%d]: %c", idx, stack.items[idx])
	}
}
