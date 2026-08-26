package queue

import (
	"sync"
	"testing"
	"time"
)

func TestQueuePushPop(t *testing.T) {
	t.Parallel()

	q := NewQueue(2)
	if !q.Push(1, time.Second) || !q.Push(2, time.Second) {
		t.Fatal("Push() failed")
	}
	if got := q.Pop(time.Second); got != 1 {
		t.Fatalf("first Pop() = %v, want 1", got)
	}
	if got := q.Pop(time.Second); got != 2 {
		t.Fatalf("second Pop() = %v, want 2", got)
	}
}

func TestQueueTimeout(t *testing.T) {
	t.Parallel()

	q := NewQueue(1)
	if !q.Push(1, time.Second) {
		t.Fatal("initial Push() failed")
	}
	if q.Push(2, time.Millisecond) {
		t.Fatal("Push() to full queue succeeded")
	}

	empty := NewQueue(1)
	if got := empty.Pop(time.Millisecond); got != nil {
		t.Fatalf("Pop() from empty queue = %v, want nil", got)
	}
}

func TestQueueCloseIsConcurrentAndIdempotent(t *testing.T) {
	t.Parallel()

	for range 100 {
		q := NewQueue(1)
		var wait sync.WaitGroup
		wait.Add(3)
		go func() {
			defer wait.Done()
			q.Push(1, time.Second)
		}()
		go func() {
			defer wait.Done()
			q.Close()
		}()
		go func() {
			defer wait.Done()
			q.Close()
		}()
		wait.Wait()
		if q.Push(2, time.Second) {
			t.Fatal("Push() after Close() succeeded")
		}
		if got := q.Pop(time.Second); got != nil {
			t.Fatalf("Pop() after Close() = %v, want nil", got)
		}
	}
}

func TestQueueCloseUnblocksWaiters(t *testing.T) {
	t.Parallel()

	full := NewQueue(1)
	if !full.Push(1, time.Second) {
		t.Fatal("initial Push() failed")
	}
	pushResult := make(chan bool, 1)
	go func() {
		pushResult <- full.Push(2, time.Hour)
	}()
	full.Close()
	if <-pushResult {
		t.Fatal("blocked Push() succeeded after Close()")
	}

	empty := NewQueue(1)
	popResult := make(chan interface{}, 1)
	go func() {
		popResult <- empty.Pop(time.Hour)
	}()
	empty.Close()
	if got := <-popResult; got != nil {
		t.Fatalf("blocked Pop() = %v, want nil", got)
	}
}
