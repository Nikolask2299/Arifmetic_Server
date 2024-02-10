package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

type UserTask struct {
	Id string 
	task string
	URL string
}

func NewUserTask(request *http.Request) ([]*UserTask, error) {
	res := make([]*UserTask, 0, 10)
	if request.Header.Get("Content-Type") == "application/json" {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		
		var buff []string

		err = json.Unmarshal(body, &buff)
		if err != nil {
			return nil, err
		}

		url := request.URL.Path
		id := NewId()

		for _, tas := range buff {
			res = append(res, &UserTask{Id: id, task: tas, URL: url})
		}
	} else if request.Header.Get("Content-Type") == "text/plain" {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		url := request.URL.Path
		id := NewId()
		res = append(res, &UserTask{Id: id, task: string(body), URL: url})
	}
	return res, nil
}

func (a *AgentServiceInput)MainOrchestrator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		task, err := NewUserTask(r)
		if err != nil {
			fmt.Fprint(w, "Error creating incorrect user task")
		} else {
			fmt.Fprintln(w, task[0].Id)
			for _, ts := range task {
				gf := ts
				go func(gf *UserTask) {
					a.Push(*gf)
				}(gf) 
			}			
		}
	} else if r.Method == "GET" {

	} else {
		fmt.Fprintln(w, "Invalid method")
	}
}

func NewId() string {
	var res string
	for i := 0; i < 3; i++ {
		num := rand.Intn(100)
		res += strconv.Itoa(num)
	}	
	return res
}