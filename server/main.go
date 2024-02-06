package main

import (
	"fmt"
	"pkg/config"
)

func main() {
	cfg := config.Mustload()
	fmt.Println(cfg)
}