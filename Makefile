build:
	go build -o apiserver . && mv apiserver ./bin

fmt:
	go fmt ./...