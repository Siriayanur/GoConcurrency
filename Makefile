build:
	go build -o bin/main main.go wire_gen.go

run:
	go run main.go wire_gen.go

all:
	build run