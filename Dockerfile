FROM golang:1.11.2-alpine3.8
WORKDIR /go
ADD main.go .
RUN ["go", "build", "main.go"]
ENTRYPOINT ["./main"]