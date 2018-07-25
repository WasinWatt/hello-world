TAG=$(shell git rev-list HEAD --max-count=1 --abbrev-commit)

dev:
	go run server.go

build:
	go build -ldflags "-X main.version=$(TAG)" -o hello-world .

pack: build
	docker build -t gcr.io/game-bot-b2fe9/hello-world:$(TAG) .