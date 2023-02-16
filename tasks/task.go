package tasks

import (
	"runtime"
	"sync"
)

type Task interface {
	AddTask(task func())
	Wait()
}

type task struct {
	sync.Mutex
	maxConcurrentTask int
	tasks             []func()
	wg                sync.WaitGroup
	c                 chan struct{}
}

func NewTask(cfg *Config) Task {
	var t *task = &task{
		maxConcurrentTask: runtime.NumCPU(),
		tasks:             []func(){},
	}

	if cfg != nil {
		if cfg.MaxConcurrentTask > 0 {
			t.maxConcurrentTask = cfg.MaxConcurrentTask
		}
	}

	t.c = make(chan struct{}, t.maxConcurrentTask)

	return t
}

func (t *task) AddTask(task func()) {
	t.Lock()
	t.tasks = append(t.tasks, task)
	t.Unlock()
}

func (t *task) Wait() {
	for i := 0; i < len(t.tasks); i++ {
		t.wg.Add(1)

		go func(taskIndex int) {
			defer func() {
				t.wg.Done()
				<-t.c
			}()

			t.c <- struct{}{}
			t.tasks[taskIndex]()
		}(i)
	}

	t.wg.Wait()
}
