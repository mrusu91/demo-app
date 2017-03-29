FROM golang:1.6

ENV USER root

WORKDIR /go/src/github.com/viruxel/demo-app
COPY . .

ENTRYPOINT ["/usr/local/go/bin/go", "run", "main.go"]
