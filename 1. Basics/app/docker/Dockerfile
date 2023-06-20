#FROM golang:alpine
FROM alpine

RUN apk add --no-cache go

WORKDIR /opt/app

CMD ["go", "run", "hello-world.go"]