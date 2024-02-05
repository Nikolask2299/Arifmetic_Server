package config

import "time"

type Config struct {
	server string `json:"server"`
	port int 	`json:"port"`
	timeout time.Duration `json:`
	storagePath string
	CountAgent int
}

