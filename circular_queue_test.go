package go_circular_queue

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	queue := NewCircularQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	if queue.IsFull() {
		fmt.Println("Queue is full.")
	}
	v1, ok := queue.Dequeue()
	v2, ok := queue.Dequeue()
	v3, ok := queue.Dequeue()
	if ok {
		fmt.Println("Values %d, %d, %d", v1, v2, v3)
	}
	if queue.IsEmpty() {
		fmt.Println("Queue is empty.")
	}
}

func TestNewCircularQueue(t *testing.T) {
	size := 100
	queue := NewCircularQueue(size)
	if queue == nil {
		t.Errorf("NewCircularQueue returned nil")
	}
	if queue.Size() != size {
		t.Errorf("Expected queue size of %d but got %d", size, queue.Size())
	}

	for i := 0; i < size; i++ {
		queue.Enqueue(i)
	}

	if !queue.IsFull() {
		t.Errorf("Expected the queue to be full.")
	}

	if queue.Enqueue(101) {
		t.Errorf("Expected not enqueue extra value.")
	}

	for i := 0; i < size; i++ {
		if v, ok := queue.Dequeue(); v != i || !ok {
			t.Errorf("Expected return value of %d got %d with return status %t.", i, v, ok)
		}
	}

	if !queue.IsEmpty() {
		t.Errorf("Expected queue to be empty.")
	}

	if _, ok := queue.Dequeue(); ok {
		t.Errorf("Expected to return status false but received true")
	}
}

func BenchmarkNewCircularQueue(b *testing.B) {
	queue := NewCircularQueue(b.N)
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
		queue.IsFull()
		queue.IsEmpty()
	}
	for i:= 0; i < b.N; i++ {
		queue.Dequeue()
		queue.IsEmpty()
		queue.IsFull()
	}
}