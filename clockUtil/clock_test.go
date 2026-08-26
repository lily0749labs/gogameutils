package clockUtil

import (
	"testing"
	"time"
)

func TestClockExecutesRequestedCount(t *testing.T) {
	timer := CreateClock()
	t.Cleanup(timer.Close)

	executed := make(chan int, 2)
	job, ok := timer.AddClock(5*time.Millisecond, 1, 2, func(param any) {
		executed <- param.(int)
	}, 42)
	if !ok || job == nil {
		t.Fatal("AddClock() did not add a job")
	}

	deadline := time.After(time.Second)
	for i := 0; i < 2; i++ {
		select {
		case got := <-executed:
			if got != 42 {
				t.Fatalf("callback parameter = %d, want 42", got)
			}
		case <-deadline:
			t.Fatalf("timed out after %d callback(s)", i)
		}
	}
}

func TestClockRejectsJobsAfterClose(t *testing.T) {
	timer := CreateClock()
	timer.Close()

	if job, ok := timer.AddClock(time.Millisecond, 1, 1, func(any) {}, nil); ok || job != nil {
		t.Fatalf("AddClock() after Close() = (%v, %t), want (nil, false)", job, ok)
	}
}
