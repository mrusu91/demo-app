FROM golang:1.6

ENV USER root

WORKDIR /go/src/github.com/viruxel/demo-app
COPY . .

RUN go get -u github.com/lib/pq

ENTRYPOINT ["/usr/local/go/bin/go", "run", "main.go"]
