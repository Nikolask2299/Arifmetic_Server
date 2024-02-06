.DEFAULT_GLOBAL := gen

running:
	go run server\main.go --config=config_file\configuration.json
.PHONY:running