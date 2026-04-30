FROM golang:1.25-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o app ./cmd/api

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/app .

EXPOSE 8080
CMD ["./app"]
