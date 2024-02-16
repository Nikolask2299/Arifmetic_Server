package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"services/pkg/agent"
	"services/pkg/config"
	"time"
)

func HandleHome(rw http.ResponseWriter, r *http.Request) {
	path := filepath.Join("E:/", "prim", "Arifmetic_Server", "html", "site.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		http.Error(rw, err.Error(), 400)
		return
	}
}



func main() {
	cfg := config.Mustload()
	timeout, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		fmt.Println("Timeout is not a valid duration")
	}

	agents := agent.NewAgentServiceInput(timeout)
	agent.NewCountDemon(cfg.CountAgent, agents)

	
	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/" + cfg.Server, agents.MainOrchestrator)
	http.ListenAndServe(":" + cfg.Port, nil)
}