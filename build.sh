#!bin/sh

export GO111MODULE="on"

go get .

go build -o ./bin/valo_tracker.exe ./main.go