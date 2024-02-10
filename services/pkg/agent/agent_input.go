package agent

import (
	"errors"
	"fmt"
	"services/pkg/arifm"
	"sync"
	"time"
)

type AgentServiceInput struct {
	ChanInputTask chan *UserTask
	mux sync.RWMutex
	timeout time.Duration
}

func NewCountDemon(count int, agent *AgentServiceInput) {
	for i := 0; i < count; i++ {
		go Demon(agent)
	}
}

func Demon(agent *AgentServiceInput) {
	for {
		task := agent.GetTask()
		if task == nil {
			time.Sleep(time.Second * 10)
			continue
		}

		res, err := arifm.ArifmeticServer(task.task)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}

func (a *AgentServiceInput) GetTask() *UserTask {
	a.mux.RLock()
	defer a.mux.RUnlock()
	select {
		case tsk := <- a.ChanInputTask:
			return tsk
		case <-time.After(2 * time.Second):
			return nil
	}
}

func NewAgentServiceInput(timeout time.Duration) *AgentServiceInput {
	return &AgentServiceInput{ChanInputTask: make(chan *UserTask, 10), mux: sync.RWMutex{}, timeout: timeout}
}

func (a *AgentServiceInput) Push(task UserTask) error {
	a.mux.RLock()
	defer a.mux.RUnlock()
	select {
		case a.ChanInputTask <- &task:
			return nil
		case <-time.After(3 * time.Second):
			return errors.New("AgentService is unavailable for push")
	}
}