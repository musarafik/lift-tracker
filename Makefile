build:
	go build -o ./ ./src/...

run:
	./cmd

clean: 
	rm cmd

all: clean build run 