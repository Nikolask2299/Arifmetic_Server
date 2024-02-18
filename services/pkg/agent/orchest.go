package agent

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

type OutAnswer struct {
	Id int `json:"id"`
	Task   string `json:"task"`
	Answer int `json:"answer"`
	Status string `json:"status"`
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
			res = append(res, &UserTask{id: id, task: tas, URL: url})
		}
	} else if request.Header.Get("Content-Type") == "text/plain" {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		url := request.URL.Path
		id := NewId()
		res = append(res, &UserTask{id: id, task: string(body), URL: url})
	}
	return res, nil
}

func (a *MainOrchestratorService)MainOrchestrator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		task, err := NewUserTask(r)
		if err != nil {
			fmt.Fprint(w, "Error creating incorrect user task")
		} else {
			index := make([]int, 0, len(task))
			for _, ts := range task {
				gf := ts
				a.dataout.dataBool[gf.id] = true
				index = append(index, gf.id)
				go func(gf *UserTask) {
					a.AgentInp.Push(*gf)
				}(gf) 
			}	

			mainIndex := NewId()
			a.dataout.dataindex[mainIndex] = index
			fmt.Fprintln(w, mainIndex)
		}
	} else if r.Method == "GET" {
		idst := r.Header.Get("id")
		id, _ := strconv.Atoi(idst)
		masout := make([]OutAnswer, 0, 10)
		if vl, ok := a.dataout.dataindex[id]; ok {
			for _, ts := range vl {
				answ := a.GetAnswerData(ts)
				if answ != nil {
					outwansw1 := OutAnswer{
						Id: answ.Id,
						Task: answ.task,
						Answer: answ.answer,
						Status: "OK",
					}
					delete(a.dataout.dataBool, answ.Id)
					masout = append(masout, outwansw1)
				} else {
					if !a.dataout.dataBool[id] {
						continue
					}
					outwansw2 := OutAnswer{
						Id: ts,
						Task: "NULL",
						Answer: 0,
						Status: "WORKING",
					}
					masout = append(masout, outwansw2)
				}	
			}
			body, _ := json.Marshal(masout)
			fmt.Fprintln(w, string(body))
		} else {
			outwansw := OutAnswer{
				Id: id,
				Task: "NULL",
				Answer: 0,
				Status: "NOT FOUND",
			}
			masout = append(masout, outwansw)
			body, _ := json.Marshal(masout)
			fmt.Fprintln(w, string(body))
		}
	} else {
		fmt.Fprintln(w, "Invalid method")
	}
}

func NewId() int {
	var res int
	for i := 0; i < 3; i++ {
		num := rand.Intn(100)
		res += num
	}	
	return res
}