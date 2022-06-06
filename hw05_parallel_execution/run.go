package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ch := make(chan interface{}, n-1)
	quit := make(chan interface{}, m)
	var wg sync.WaitGroup
	var currentErrors int
	i := 0
METKA:
	for {
		select {
		case <-quit:
			currentErrors++
			if currentErrors >= m {
				wg.Wait()
				return ErrErrorsLimitExceeded
			}
		default:
			ch <- nil
			wg.Add(1)
			go func(task Task) {
				if task() != nil {
					quit <- nil
				}
				<-ch
				wg.Done()
			}(tasks[i])
			i++
			if i >= len(tasks) {
				wg.Wait()
				break METKA
			}
		}
	}
	return nil
}
