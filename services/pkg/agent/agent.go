package agent

import (
	"errors"
	"services/pkg/orchestrator"
	"sync"
	"time"
)

type AgentService struct {
	ChanResponseChan chan orchestrator.UserTask
	mux sync.RWMutex
}

func NewAgentService() *AgentService {
	return &AgentService{ChanResponseChan: make(chan orchestrator.UserTask, 10), mux: sync.RWMutex{}}
}

func (a *AgentService) Push(task orchestrator.UserTask) error {
	select {
		case a.ChanResponseChan <- task:
			return nil
		case <-time.After(time.Second):
			return errors.New("AgentService is unavailable for push")
	}
}