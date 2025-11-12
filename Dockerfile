# builder
FROM golang:1.24-alpine AS builder
WORKDIR /src
COPY go.mod go.sum ./
RUN apk add --no-cache git
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/effectivetask ./cmd

# runtime
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/bin/effectivetask /usr/local/bin/effectivetask
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/effectivetask"]