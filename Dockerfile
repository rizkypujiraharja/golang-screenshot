# Build stage
FROM golang:1.20-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /app
COPY . .

# Fetch dependencies.
# RUN go get -d -v
RUN go mod download
RUN go mod verify

RUN go build main.go

# Run stage
FROM chromedp/headless-shell:latest
# FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /go/bin/golang-screenshot
EXPOSE 8080


ENTRYPOINT ["/go/bin/golang-screenshot"]
