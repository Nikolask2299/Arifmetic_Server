package main

import (
	"fmt"
	"net/http"
	"services/pkg/agent"
	"services/pkg/config"
	"time"
)

func main() {
	cfg := config.Mustload()
	timeout, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		fmt.Println("Timeout is not a valid duration")
	}
	agents := agent.NewAgentService(timeout)
	agent.NewCountDemon(cfg.CountAgent, agents)

	http.HandleFunc("/" + cfg.Server, agents.MainOrchestrator)
	http.ListenAndServe(":" + cfg.Port, nil)
}