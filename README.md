## Go Circular-Queue

Circular Queue is a linear data structure in which the operations are performed based on FIFO (First In First Out) principle and the last position is connected back to the first position to make a circle.
It provides a high-performance, circular queue (ring buffer) implementation in golang.
## Getting Start

```bash
go get -v github.com/arazmj/go-circular-queue
```

```go
import (
 "github.com/arazmj/go-circular-queue"
)

func main() {
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
```

## Contrib

Welcome


## LICENSE

BSD