FROM golang:1.21.0-alpine

WORKDIR /backend

COPY go.mod go.sum /backend/

RUN go mod download

EXPOSE 4000
CMD ["go","run","main.go"]