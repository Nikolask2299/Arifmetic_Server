package agent

import (
	"fmt"
	"net/http"
)

type UserTask struct {
	respWrite http.ResponseWriter
	request *http.Request
}


func NewUserTask(request *http.Request, responseWriter http.ResponseWriter) *UserTask {
	return &UserTask{respWrite: responseWriter, request: request}
}

func (a *AgentService)MainOrchestrator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		task := NewUserTask(r, w)
		a.Push(*task)
	} else if r.Method == "GET" {
			
	} else {
		fmt.Fprintln(w, "Invalid method")
	}
}
