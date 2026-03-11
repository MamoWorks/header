FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o header main.go

FROM alpine:3.19
COPY --from=builder /app/header /header
EXPOSE 8899
ENTRYPOINT ["/header"]
