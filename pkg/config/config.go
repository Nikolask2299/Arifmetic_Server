package config

import "time"

type Config struct {
	server string `json:"server"`
	port int 	`json:"port"`
	timeout time.Duration `json:"timeout"`
	storagePath string `json:"storage_path"`
	CountAgent int `json:"count_server_agent"`
}

 func Mustload() Config {


	
 }