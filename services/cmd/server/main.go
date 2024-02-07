package main

import (
	"fmt"
	"net/http"
	"services/pkg/config"
	"services/pkg/orchestrator"
	"time"
)

func main() {
	cfg := config.Mustload()
	timeout, err := time.ParseDuration(cfg.Timeout)
	if err != nil {
		fmt.Println("Timeout is not a valid duration")
	}
	

	http.HandleFunc("/" + cfg.Server, orchestrator.MainOrchestrator)
	http.ListenAndServe(":" + cfg.Port, nil)
}