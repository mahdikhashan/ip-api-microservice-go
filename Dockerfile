# syntax=docker/dockerfile:1
ARG GO_VERSION=1.23.4

FROM golang:${GO_VERSION}-alpine AS builder
LABEL authors="mahdikhashan"

RUN apk update && apk add --no-cache alpine-sdk git

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /service-ip

---

FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /service-ip /service-ip

EXPOSE 8080

CMD [ "/service-ip" ]
