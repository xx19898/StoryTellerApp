FROM golang:1.21.0-alpine
WORKDIR /backend
COPY go.mod go.sum /backend/
RUN go mod download &&