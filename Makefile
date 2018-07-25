TAG=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

dev:
	go run server.go

build:
	go build -ldflags "-X main.version=$(TAG)" -o server .