package circlequeue

// CircularBuffer 是循环缓存的结构体
type CircularBuffer struct {
	size  int
	data  []int
	front int
	rear  int
}

// NewCircularBuffer 创建一个新的循环缓存
func NewCircularBuffer(size int) *CircularBuffer {
	return &CircularBuffer{
		size:  size,
		data:  make([]int, size),
		front: 0,
		rear:  0,
	}
}

// Enqueue 将数据入队
func (cb *CircularBuffer) Enqueue(item int) {
	if cb.isFull() {
		cb.front = (cb.front + 1) % cb.size // 循环覆盖最旧的数据
	}
	cb.data[cb.rear] = item
	cb.rear = (cb.rear + 1) % cb.size

}

// Dequeue 将数据出队
func (cb *CircularBuffer) Dequeue() int {
	if cb.isEmpty() {
		return -1 // 表示空队列
	}
	item := cb.data[cb.front]
	cb.front = (cb.front + 1) % cb.size
	return item
}

// isFull 检查队列是否已满
func (cb *CircularBuffer) isFull() bool {
	return (cb.rear+1)%cb.size == cb.front
}

// isEmpty 检查队列是否为空
func (cb *CircularBuffer) isEmpty() bool {
	return cb.front == cb.rear
}

func (cb *CircularBuffer) QueueLength() int {
	return (cb.rear - cb.front + cb.size) % cb.size
}
