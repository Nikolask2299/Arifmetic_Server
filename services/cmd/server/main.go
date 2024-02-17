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
	path := filepath.Join("html", "site.html")

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		fmt.Println(err)
		http.Error(rw, err.Error(), 400)
		return
	}

	err = tmpl.Execute(rw, nil)
	if err != nil {
		fmt.Println(err)
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

	agentsInp := agent.NewAgentServiceInput(timeout)
	agentsOut := agent.NewAgentServiceOutput(timeout)

	mainOrcest := agent.NewMainOrchestratorService(agentsInp, agentsOut)

	agent.NewCountDemon(cfg.CountAgent, mainOrcest.AgentInp, mainOrcest.AgentOut)
	go mainOrcest.Output()
	fmt.Println("Server OK")
	http.HandleFunc("/", HandleHome)
	http.HandleFunc("/" + cfg.Server, mainOrcest.MainOrchestrator)
	http.ListenAndServe(":" + cfg.Port, nil)
}