.PHONY: run

export GO111MODULE=on

httpserver:
	go run server/cmd/cli/main.go

subscriber:
	go run messaging/subscriber.go
