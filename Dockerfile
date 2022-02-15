FROM golang:1.17-alpine

RUN apk add --no-cache musl-dev gcc make g++ file
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
RUN mkdir -p /app
COPY . /app
WORKDIR /app

ENV GO111MODULE=on
ENV APP_MODE=dev
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64
EXPOSE 8000