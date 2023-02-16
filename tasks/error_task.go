package tasks

import (
	"runtime"
	"sync"
)

type ErrorTask interface {
	AddTask(task func() error)
	Wait() error
}

type errorTask struct {
	sync.Mutex
	maxConcurrentTask int
	tasks             []func() error
	wg                sync.WaitGroup
	c                 chan struct{}
	errsChan          chan error
}

func NewErrorTask(cfg *Config) ErrorTask {
	var et *errorTask = &errorTask{
		maxConcurrentTask: runtime.NumCPU(),
		tasks:             []func() error{},
	}

	if cfg != nil {
		if cfg.MaxConcurrentTask > 0 {
			et.maxConcurrentTask = cfg.MaxConcurrentTask
		}
	}

	et.c = make(chan struct{}, et.maxConcurrentTask)

	return et
}

func (et *errorTask) AddTask(task func() error) {
	et.Lock()
	et.tasks = append(et.tasks, task)
	et.Unlock()
}

func (et *errorTask) Wait() error {
	var err error

	et.errsChan = make(chan error, len(et.tasks))

	for i := 0; i < len(et.tasks); i++ {
		et.wg.Add(1)

		go func(taskIndex int) {
			var errRoutine error

			defer func() {
				et.wg.Done()
				<-et.c
			}()

			et.c <- struct{}{}

			if len(et.errsChan) == 0 {
				errRoutine = et.tasks[taskIndex]()

				if errRoutine != nil {
					et.errsChan <- errRoutine
				}
			}
		}(i)
	}

	et.wg.Wait()

	if len(et.errsChan) > 0 {
		err = <-et.errsChan

		return err
	}

	return nil
}
