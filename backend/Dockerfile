# syntax=docker/dockerfile:1

FROM golang:latest AS build-stage

WORKDIR ../backend

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /rtag

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /rtag /rtag

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/rtag"]