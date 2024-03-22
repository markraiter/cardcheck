########-- Build stage --########

FROM golang:1.22-alpine AS builder
LABEL authors="MarkRaiter"

WORKDIR /opt

COPY . /opt

RUN go build -o ./runner ./cmd

########-- Deploy stage --########

FROM alpine:3.18
LABEL authors="MarkRaiter"

WORKDIR /opt 

COPY --from=builder /opt/runner /opt/

EXPOSE 5555

ENTRYPOINT /opt/runner