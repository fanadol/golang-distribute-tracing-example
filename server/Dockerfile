# Builder Image
FROM golang AS builder
WORKDIR /usr/src/app

ENV GO111MODULE=on

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /usr/src/app/bin/main /usr/src/app/cmd/cli/main.go

# Main Image
FROM alpine:latest
WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/bin/main .

EXPOSE 8080

CMD ["./main"]
