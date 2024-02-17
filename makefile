.DEFAULT_GLOBAL := gen

running:
	go run services/cmd/server/main.go --config=config_file/configuration.json 
.PHONY:running