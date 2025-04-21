.PHONY: test

build:
	go build -o dist/game cmd/gogame/gogame.go

run:
	go run cmd/gogame/gogame.go

test:
	go test -v ./test -coverpkg ./...
