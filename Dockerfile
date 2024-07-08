FROM golang:1.22.4 AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8000

ENTRYPOINT ["/app/app"]