FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:3.23

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/config ./config

EXPOSE 9001

ENTRYPOINT ["./app"]

# CMD ["./app", "server"]
