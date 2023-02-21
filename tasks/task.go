package tasks

import (
	"runtime"
	"sync"
)

type Task interface {
	Go(task func())
	Wait()
}

type task struct {
	maxConcurrentTask int
	wg                sync.WaitGroup
	c                 chan struct{}
}

func NewTask(cfg *Config) Task {
	var t *task = &task{
		maxConcurrentTask: runtime.NumCPU(),
	}

	if cfg != nil {
		if cfg.MaxConcurrentTask > 0 {
			t.maxConcurrentTask = cfg.MaxConcurrentTask
		}
	}

	t.c = make(chan struct{}, t.maxConcurrentTask)

	return t
}

func (t *task) Go(task func()) {
	t.wg.Add(1)

	go func(taskToDo func()) {
		defer func() {
			t.wg.Done()
			<-t.c
		}()

		t.c <- struct{}{}
		taskToDo()
	}(task)
}

func (t *task) Wait() {
	t.wg.Wait()
}
