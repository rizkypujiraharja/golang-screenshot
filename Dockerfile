# Build stage
FROM golang:1.20-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /app
COPY . .

# Fetch dependencies.
# RUN go get -d -v
RUN go mod download
RUN go mod verify

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/golang-screenshot .

# Run stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/bin/golang-screenshot /go/bin/golang-screenshot
EXPOSE 8080

ENTRYPOINT ["/go/bin/golang-screenshot"]
