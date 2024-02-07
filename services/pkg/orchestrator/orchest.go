package orchestrator

import (
	"fmt"
	"net/http"
)


type UserTask struct {
	RespWrite http.ResponseWriter
	request *http.Request
}

func MainOrchestrator(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		
	} else if r.Method == "GET" {

	} else {
		fmt.Fprintln(w, "Invalid method")
	}
}