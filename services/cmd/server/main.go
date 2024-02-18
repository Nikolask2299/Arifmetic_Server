package main

import (
	"fmt"
	"html/template"
	"net/http"
	"services/pkg/agent"
	"services/pkg/config"
	"time"
)

type HTML struct {
	path string
}

func (p *HTML)HandleHome(rw http.ResponseWriter, r *http.Request) {
	
	tmpl, err := template.ParseFiles(p.path)
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

	html := &HTML{
		path: cfg.HTMLpath,
	}

	agent.NewCountDemon(cfg.CountAgent, mainOrcest)
	go mainOrcest.Output()
	fmt.Println("Server OK")

	http.HandleFunc("/", html.HandleHome)
	http.HandleFunc("/" + cfg.Server, mainOrcest.MainOrchestrator)
	http.ListenAndServe(":" + cfg.Port, nil)
}