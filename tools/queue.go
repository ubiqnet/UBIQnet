package tools

import (
	"sync"
	"time"
)

type QueueItem = interface{}

type QueueConsumer func(QueueItem)

type Queue struct {
	sync.RWMutex

	items []QueueItem

	count int

	running bool
}

func (q *Queue) Enqueue(v ...QueueItem) {
	q.Lock()
	q.items = append(q.items, v...)
	q.Unlock()
}

func (q *Queue) Dequeue() QueueItem {
	q.Lock()
	defer q.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	v := q.items[0]

	q.items = q.items[1:]

	return v
}

func (q *Queue) Size() int {
	q.RLock()
	defer q.RUnlock()

	return len(q.items)
}

func (q *Queue) Done() {
	q.Lock()
	defer q.Unlock()

	if q.count > 0 {
		q.count--
	}
}

func (q *Queue) Run(consumer QueueConsumer) {
	if q.running {
		return
	}

	q.running = true

	ticker := time.NewTicker(time.Millisecond)
	go func() {

		for range ticker.C {
			q.worker(consumer)
		}
	}()

}

func (q *Queue) worker(consumer QueueConsumer) {
	q.RLock()

	q.RUnlock()
	if v := q.Dequeue(); v != nil {
		q.Lock()
		q.count++
		q.Unlock()
		consumer(v)
	}
}

func NewQueue() *Queue {
	return &Queue{
		running: false,
	}
}
