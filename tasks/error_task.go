package tasks

import (
	"runtime"
	"sync"
)

type ErrorTask interface {
	Go(task func() error)
	Wait() error
}

type errorTask struct {
	maxConcurrentTask int
	wg                sync.WaitGroup
	c                 chan struct{}
	errsChan          chan error
}

func NewErrorTask(cfg *Config) ErrorTask {
	var et *errorTask = &errorTask{
		maxConcurrentTask: runtime.NumCPU(),
	}

	if cfg != nil {
		if cfg.MaxConcurrentTask > 0 {
			et.maxConcurrentTask = cfg.MaxConcurrentTask
		}
	}

	et.c = make(chan struct{}, et.maxConcurrentTask)
	et.errsChan = make(chan error, 1)

	return et
}

func (et *errorTask) Go(task func() error) {
	et.wg.Add(1)

	go func(taskToDo func() error) {
		var errRoutine error

		defer func() {
			et.wg.Done()
			<-et.c
		}()

		et.c <- struct{}{}

		if len(et.errsChan) == 0 {
			errRoutine = taskToDo()

			if errRoutine != nil {
				et.errsChan <- errRoutine
			}
		}
	}(task)
}

func (et *errorTask) Wait() error {
	var err error

	et.wg.Wait()

	if len(et.errsChan) > 0 {
		err = <-et.errsChan

		return err
	}

	return nil
}
