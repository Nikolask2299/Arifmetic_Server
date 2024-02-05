package config

import "time"

type Config struct {
	server string `json:"server"`
	port int 	
	timeout time.Duration
	storagePath string
	CountAgent int
}