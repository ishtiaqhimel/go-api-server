build:
	go build -o apiserver . && mkdir -p ./bin && mv apiserver ./bin

fmt:
	go fmt ./...