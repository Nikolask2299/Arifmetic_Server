package agent

import (
	"errors"
	"sync"
	"time"
)

type UserAnswer struct {
	Id     int
	URL    string
	task   string
	answer int
}

type AgentServiceOutput struct {
	ChanOutputAnswer chan *UserAnswer
	mux              sync.RWMutex
	timeout time.Duration
}

func NewUserAnswer(task UserTask, answer int) *UserAnswer {
	return &UserAnswer{
		Id: task.id,
		URL: task.URL,
		task: task.task,
		answer: answer,
	}
}

func NewAgentServiceOutput(timeout time.Duration) *AgentServiceOutput {
	return &AgentServiceOutput{
		ChanOutputAnswer: make(chan *UserAnswer, 10),
		mux: sync.RWMutex{},
		timeout: timeout,
	}
}

func(a *AgentServiceOutput) GetAnswer() *UserAnswer {
	a.mux.RLock()
	defer a.mux.RUnlock()
	select {
		case ans := <- a.ChanOutputAnswer:
			return ans
		case <-time.After(3 * time.Second):
			return nil
	}
}

func (a *AgentServiceOutput) PushAnswer(usr UserAnswer) error {
	a.mux.RLock()
	defer a.mux.RUnlock()
	
	select {
		case a.ChanOutputAnswer <- &usr:
			return nil
		case <- time.After(3 * time.Second):
			return errors.New("AgentService is unavailable for push")
	}	
}

