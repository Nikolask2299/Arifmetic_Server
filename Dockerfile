FROM golang:1.21-bookworm

COPY ./ ./
RUN go build -o services\cmd\server\main.go --config=config_file\configuration.json

CMD [ "./makefile" ]