# syntax=docker/dockerfile:1
FROM golang:1.16-alpine AS build

RUN apk add --update git;

WORKDIR ${GOPATH}/app/

COPY *.go ./
COPY go.mod ./
COPY go.sum ./

RUN go build -o app . 

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=build /go/app/app .

CMD ["/app/app"]