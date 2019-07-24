FROM golang:latest

WORKDIR /hubb
ADD . /hubb

RUN go install github.com/pilu/fresh && \ 
    go generate ./...

ENV PORT 8080
EXPOSE 8080

CMD ["go", "run", "main.go"]
