package queue

import (
	"sync"
)

type FIFOQueue interface {
	Add(item interface{})
	Len() int
	Pop() (item interface{}, shutdown bool)
	Done()
	ShutDown()
	HasShutDown() bool
}

func NewQueue() FIFOQueue {
	return &Queue{
		cond:     sync.NewCond(&sync.Mutex{}),
		data:     make([]interface{}, 0),
		shutDown: false,
	}
}

type Queue struct {
	cond     *sync.Cond
	data     []interface{}
	shutDown bool
}

func (q *Queue) Add(item interface{}) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if q.shutDown {
		return
	}
	q.data = append(q.data, item)
	q.cond.Signal()
}

func (q *Queue) Len() int {
	return len(q.data)
}

// Pop blocks until it can return an item to be processed. If shutdown = true,
// the caller should end their goroutine. You must call Done with item when you
// have finished processing it.
func (q *Queue) Pop() (item interface{}, shutdown bool) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	if q.shutDown {
		return nil, true
	}

	for q.Len() == 0 && !q.HasShutDown() {
		q.cond.Wait()
	}

	item = q.data[0]
	q.data = q.data[1:]
	return item, false
}

func (q *Queue) Done() {

}

func (q *Queue) ShutDown() {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.data = nil
	q.shutDown = true

}

func (q *Queue) HasShutDown() bool {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	return q.shutDown
}
