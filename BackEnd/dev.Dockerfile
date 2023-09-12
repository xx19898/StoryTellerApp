FROM golang:1.21.0-alpine

WORKDIR /backend

COPY go.mod .
COPY go.sum .

RUN go mod download

EXPOSE 4000
CMD ["go","run","main.go"]