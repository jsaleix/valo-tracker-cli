.PHONY: all test clean

install:
	go get .
build:
	go build -o ./bin/valo_tracker.exe ./main.go