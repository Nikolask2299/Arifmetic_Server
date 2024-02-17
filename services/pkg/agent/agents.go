package agent

import (
	"fmt"
	"time"
	"services/pkg/arifm"
)

func NewCountDemon(count int, agentInp *AgentServiceInput, agentOut *AgentServiceOutput) {
	for i := 0; i < count; i++ {
		go Demon(agentInp, agentOut)
	}
}

func Demon(agentInp *AgentServiceInput, agentOut *AgentServiceOutput) {
	for {
		task := agentInp.GetTask()
		if task == nil {
			time.Sleep(time.Second * 10)
			continue
		}

		res, err := arifm.ArifmeticServer(task.task)
		if err != nil {
			fmt.Println(err)
		}
		anwer := NewUserAnswer(*task, res)
		agentOut.PushAnswer(anwer)
	}
}

func (mainOrcServ *MainOrchestratorService) Output() {
	for {
		answ := mainOrcServ.AgentOut.GetAnswer()

		if answ == nil {
			time.Sleep(3 * time.Second)
			continue
		}

		mainOrcServ.database[answ.Id] = answ
	}	
}
