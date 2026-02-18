FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app ./cmd/main.go

FROM alpine:3.19

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/app .
COPY --from=builder /app/.env .

RUN mkdir -p /app/tsv_files /app/reports /app/logs

EXPOSE 8080

CMD ["./app"]