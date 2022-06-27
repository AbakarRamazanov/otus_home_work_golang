package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (err error) {
	tool := &TaskerTool{
		WorkersCount:   n,
		MaxErrorsCount: int64(m),
		CountErr:       0,
		ch:             make(chan Task),
		wg:             new(sync.WaitGroup),
	}
	tool.createGoroutines()
	defer tool.finishGoroutines()
	for i := 0; i < len(tasks); i++ {
		if tool.checkExtendedErrors() {
			return ErrErrorsLimitExceeded
		}
		tool.ch <- tasks[i]
	}
	return nil
}

type TaskerTool struct {
	TasksCount     int
	WorkersCount   int
	MaxErrorsCount int64
	CountErr       int64
	ch             chan Task
	wg             *sync.WaitGroup
}

func (t *TaskerTool) createGoroutines() {
	for i := 0; i < t.WorkersCount; i++ {
		t.wg.Add(1)
		go t.run()
	}
}

func (t *TaskerTool) checkExtendedErrors() bool {
	return atomic.LoadInt64(&t.CountErr) >= t.MaxErrorsCount
}

func (t *TaskerTool) run() {
	defer t.wg.Done()
	for f := range t.ch {
		if f() != nil {
			atomic.AddInt64(&t.CountErr, 1)
		}
	}
}

func (t TaskerTool) finishGoroutines() {
	close(t.ch)
	t.wg.Wait()
}
