package agent

import (
	"errors"
	"fmt"
	"io"
	"services/pkg/arifm"
	"sync"
	"time"
)

type AgentService struct {
	ChanResponseChan chan UserTask
	mux sync.RWMutex
	timeout time.Duration
}

func NewCountDemon(count int, agent *AgentService) {
	for i := 0; i < count; i++ {
		go Demon(agent)
	}
}

func Demon(agent *AgentService) {
	for {
		task := agent.GetTask()
		if task.request == nil {
			time.Sleep(time.Second * 10)
			continue
		}

		body, _ := io.ReadAll(task.request.Body)
		res, err := arifm.ArifmeticServer(string(body))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(res)
	}
}


func (a *AgentService) GetTask() UserTask {
	a.mux.RLock()
	defer a.mux.RUnlock()
	select {
		case tsk := <- a.ChanResponseChan:
			return tsk
		case <-time.After(time.Second):
			return UserTask{}
	}
}

func NewAgentService(timeout time.Duration) *AgentService {
	return &AgentService{ChanResponseChan: make(chan UserTask, 10), mux: sync.RWMutex{}, timeout: timeout}
}

func (a *AgentService) Push(task UserTask) error {
	select {
		case a.ChanResponseChan <- task:
			return nil
		case <-time.After(time.Second):
			return errors.New("AgentService is unavailable for push")
	}
}