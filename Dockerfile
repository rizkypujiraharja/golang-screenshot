# Build stage
FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN go build main.go

# Run stage
FROM chromedp/headless-shell:latest
# FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main /go/bin/golang-screenshot
EXPOSE 8080


ENTRYPOINT ["/go/bin/golang-screenshot"]
