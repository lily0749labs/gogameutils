package queue

import (
	"context"
	"sync/atomic"
	"time"

	queueutil "github.com/lily0749labs/goutils/queue"
)

type Queue struct {
	queue  *queueutil.Queue[any]
	closed atomic.Bool
}

// 初始化队列的长度
func NewQueue(max_queue_len int) (dc *Queue) {
	return &Queue{
		queue: queueutil.New[any](max_queue_len),
	}
}

func (q *Queue) Push(data interface{}, waittime time.Duration) bool {
	if q.closed.Load() {
		return false
	}

	ctx, cancel := context.WithTimeout(context.Background(), waittime)
	defer cancel()
	return q.queue.Push(ctx, data) == nil
}

func (q *Queue) Pop(waittime time.Duration) (data interface{}) {
	if q.closed.Load() {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), waittime)
	defer cancel()
	data, err := q.queue.Pop(ctx)
	if err != nil {
		return nil
	}
	return data
}

func (q *Queue) Close() {
	q.closed.Store(true)
	q.queue.Close()
}
