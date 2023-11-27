package circlequeue

import (
	"testing"
)

func TestCircularBufferEnqueueDequeue(t *testing.T) {
	// 创建一个容量为 5 的循环缓存
	circularBuffer := NewCircularBuffer(5)

	// 向循环缓存中添加数据
	for i := 1; i <= 8; i++ {
		circularBuffer.Enqueue(i)
	}

	// 从循环缓存中取出并检查数据
	expected := []int{5, 6, 7, 8}
	for _, val := range expected {
		item := circularBuffer.Dequeue()
		if item != val {
			t.Errorf("Expected %d, got %d", val, item)
		}
	}

	// 检查队列是否为空
	if circularBuffer.QueueLength() != 0 {
		t.Error("Queue should be empty")
	}

	// 再次向循环缓存中添加数据
	for i := 9; i <= 12; i++ {
		circularBuffer.Enqueue(i)
	}

	// 从循环缓存中取出并检查数据
	expected = []int{9, 10, 11, 12}
	for _, val := range expected {
		item := circularBuffer.Dequeue()
		if item != val {
			t.Errorf("Expected %d, got %d", val, item)
		}
	}

	// 检查队列是否为空
	if circularBuffer.QueueLength() != 0 {
		t.Error("Queue should be empty")
	}
}
