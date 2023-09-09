FROM golang:1.21.0-alpine as initial
WORKDIR /backend
COPY go.mod go.sum /backend/
RUN go mod download
COPY ./ ./
RUN go build -o /go-executable
FROM alpine
WORKDIR /storytellerBackend
COPY --from=initial ./go-executable ./
CMD ["./go-executable"]