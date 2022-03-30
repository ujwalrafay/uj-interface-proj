build:
	go build -o bin/main main.go
run:
	go run main.go

init:
	go mod init 

tidy:
	go mod tidy