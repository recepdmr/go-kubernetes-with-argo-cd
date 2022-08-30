FROM golang:1.19.0-alpine3.16

WORKDIR /go/src/app

COPY ./main.go .

RUN go build main.go
CMD ["./main"]