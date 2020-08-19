FROM golang:1.12 AS builder
WORKDIR $GOPATH/src/github.com/thazelart/terraform-validator
ENV GO111MODULE=on
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go install

FROM alpine:latest  
LABEL maintainer="Thibault Hazelart <thazelart@gmail.com>"
COPY --from=builder /go/bin/terraform-validator /usr/local/bin
WORKDIR /data
CMD ["terraform-validator", "."]
