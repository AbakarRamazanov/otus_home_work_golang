package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ch := make(chan interface{}, n-1)
	var wg sync.WaitGroup
	var currentErrors int32 = 0
	var sumWorkTask int32 = 0
	var m32 int32 = int32(m)
	for i := 0; atomic.LoadInt32(&currentErrors) <= m32 && i < len(tasks); i++ {
		ch <- nil
		wg.Add(1)
		go func(i int) {
			if tasks[i]() != nil {
				atomic.AddInt32(&currentErrors, 1)
			}
			atomic.AddInt32(&sumWorkTask, 1)
			wg.Done()
			<-ch
		}(i)
	}
	wg.Wait()
	// Place your code here.
	if currentErrors > m32 {
		return ErrErrorsLimitExceeded
	}
	return nil
}
