# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR ../backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /rtag

EXPOSE 8080

CMD [ "/rtag" ]