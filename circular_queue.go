package go_circular_queue

type CircularQueue struct {
	length 			int
	front 			int
	rear 			int
	arr				[]interface{}
}

func NewCircularQueue(size int) *CircularQueue {
	return &CircularQueue{
		length: 0,
		front:  0,
		rear:   -1,
		arr:    make([]interface{}, size),
	}
}

func (q *CircularQueue) Enqueue(value interface{}) bool {
	if q.IsFull() {
		return false
	}
	q.rear = (q.rear + 1) % len(q.arr)
	q.arr[q.rear] = value
	q.length++
	return true
}

func (q *CircularQueue) Dequeue() (val interface{}, ok bool) {
	if q.IsEmpty() {
		return nil,false
	}
	val = q.arr[q.front]
	q.arr[q.front] = nil
	q.front = (q.front + 1) % len(q.arr)
	q.length--
	return val, true
}

func (q *CircularQueue) Front() (val interface{}, ok bool) {
	if q.IsEmpty() {
		return nil, false
	}
	return q.arr[q.front], true
}

func (q *CircularQueue) Rear() (val interface{}, ok bool)  {
	if q.IsEmpty() {
		return nil, false
	}
	return q.arr[q.rear], true
}

func (q *CircularQueue) IsEmpty() bool {
	return q.length == 0
}

func (q *CircularQueue) IsFull() bool {
	return q.length == len(q.arr)
}

func (q *CircularQueue) Size() int {
	return len(q.arr)
}