package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) (err error) {
	tool := &TaskerTool{
		WorkersCount:   n,
		MaxErrorsCount: m,
		CountErr:       new(int),
		mu:             new(sync.RWMutex),
		ch:             make(chan Task),
		wg:             new(sync.WaitGroup),
	}
	tool.createGoroutines()
	defer tool.finishGoroutines()
	for i := 0; i < len(tasks); i++ {
		tool.ch <- tasks[i]
		if tool.checkExtendedErrors() {
			return ErrErrorsLimitExceeded
		}
	}
	return nil
}

type TaskerTool struct {
	TasksCount     int `faker:"boundary_start=1, boundary_end=500"`
	WorkersCount   int `faker:"boundary_start=2, boundary_end=16"`
	MaxErrorsCount int `faker:"boundary_start=1, boundary_end=100"`
	mu             *sync.RWMutex
	CountErr       *int
	ch             chan Task
	wg             *sync.WaitGroup
}

func (t TaskerTool) createGoroutines() {
	for i := 0; i < t.WorkersCount; i++ {
		t.wg.Add(1)
		go t.run()
	}
}

func (t TaskerTool) checkExtendedErrors() bool {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return *t.CountErr >= t.MaxErrorsCount
}

func (t TaskerTool) run() {
	defer t.wg.Done()
	for f := range t.ch {
		if f() != nil {
			t.mu.Lock()
			*t.CountErr++
			t.mu.Unlock()
		}
	}
}

func (t TaskerTool) finishGoroutines() {
	close(t.ch)
	t.wg.Wait()
}
