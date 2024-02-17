package agent

import (
	"errors"
	"sync"
	"time"
)

type AgentServiceInput struct {
	ChanInputTask chan *UserTask
	mux sync.Mutex
	timeout time.Duration
}

type UserTask struct {
	id int
	task string
	URL string
}

func NewAgentServiceInput(timeout time.Duration) *AgentServiceInput {
	return &AgentServiceInput{ChanInputTask: make(chan *UserTask, 10), 
		mux: sync.Mutex{}, timeout: timeout}
}

func (a *AgentServiceInput) GetTask() *UserTask {
	a.mux.Lock()
	defer a.mux.Unlock()
	select {
		case tsk := <- a.ChanInputTask:
			return tsk
		case <-time.After(2 * time.Second):
			return nil
	}
}

func (a *AgentServiceInput) Push(task UserTask) error {
	a.mux.Lock()
	defer a.mux.Unlock()
	select {
		case a.ChanInputTask <- &task:
			return nil
		case <-time.After(3 * time.Second):
			return errors.New("AgentService is unavailable for push")
	}
}
