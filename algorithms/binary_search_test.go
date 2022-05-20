package algorithms

import (
	"log"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{3, 5, 6, 8, 10, 13, 16, 18}
	key := 1
	result := binary_search(arr, key, 0, len(arr)/2, len(arr)-1)
	log.Printf("The index of the key: %d is %d", key, result)
}
