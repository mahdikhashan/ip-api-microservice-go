# syntax=docker/dockerfile:1

FROM golang:1.23.4-alpine

LABEL authors="mahdikhashan"

WORKDIR /app

COPY go.mod go.work ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /service-ip

EXPOSE 8080

CMD [ "service-ip" ]