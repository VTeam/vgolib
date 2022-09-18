package queue

import (
	"testing"
	"time"
)

const (
	TestQueueLen int = 10
)

// func TestBasic(t *testing.T) {
// 	tests := []struct{
// 		queue *Queue
// 	}
// }

func TestAdd(t *testing.T) {
	q := NewQueue()
	defer q.ShutDown()

	for i := 0; i < TestQueueLen; i++ {
		q.Add(i)
		t.Logf("add item %+v", i)
	}
	if q.Len() != TestQueueLen {
		t.Fatal("except q len", TestQueueLen)

	}
}

// func TestPop(t *testing.T) {

// }

func TestWakeUp(t *testing.T) {
	q := NewQueue()
	defer q.ShutDown()

	ticker := time.NewTicker(time.Second)
	// _, shutdown := q.Pop()
	for {
		select {
		// case :
		// 	t.Fatal("except wait", TestQueueLen)
		case <-ticker.C:
			q.Add(time.Now())
		}
	}
}
