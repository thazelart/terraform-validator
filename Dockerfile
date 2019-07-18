FROM golang:1.12

LABEL maintainer="Thibault Hazelart <thazelart@gmail.com>"
ENV GO111MODULE=on

WORKDIR $GOPATH/src/github.com/thazelart/terraform-validator
COPY . .

RUN go get -d -v ./...
RUN go install
