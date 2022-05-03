APP=main

build: clean
	go build -o bin/${APP} main.go

run:
	go run -race main.go

clean:
	go clean