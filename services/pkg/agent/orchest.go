package agent

import (
	"fmt"
	"io"
	"net/http"
)

type UserTask struct {
	Id string 
	task string
	URL string
}


func NewUserTask(request *http.Request, responseWriter http.ResponseWriter) (*UserTask, error) {
	
	body, err := io.ReadAll()
	if err == nil {
		return nil, err
	}	

	
}

func (a *AgentService)MainOrchestrator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		task, err := NewUserTask(r, w)
		if err != nil {
			fmt.Fprint(w, "Error creating incorrect user task")
		} else {

			a.Push(*task)
		}
	} else if r.Method == "GET" {
			
	} else {
		fmt.Fprintln(w, "Invalid method")
	}
}
