FROM golang:latest

WORKDIR ./src/server

CMD ["go","run","/api/main.go"]
