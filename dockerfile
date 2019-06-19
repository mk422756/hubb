FROM golang:latest

WORKDIR /hubb
ADD . /hubb

RUN go install github.com/pilu/fresh && \ 
    go generate ./...

CMD ["go", "run", "server/main.go"]