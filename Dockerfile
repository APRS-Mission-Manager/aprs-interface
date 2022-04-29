# syntax=docker/dockerfile:1

FROM golang:1.18

WORKDIR /app

# copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main ./cmd/interface/main.go

ENV APRS_USERNAME KN4CDD
ENV APRS_PASSWORD -1
CMD ["./main"]