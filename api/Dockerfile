FROM golang:1.18-alpine as builder


RUN apk update && apk add --no-cache git curl
RUN go install github.com/cosmtrek/air@v1.29.0 && chmod +x ${GOPATH}/bin/air


ENV GOPRIVATE="github.com/renasami/advents-2022-myjlab/api"
WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o /main

FROM alpine:3.9

COPY --from=builder /main .

ENTRYPOINT ["/main"]