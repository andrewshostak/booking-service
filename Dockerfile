FROM golang:1.16.7-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN go build

FROM alpine
WORKDIR /app
COPY --from=builder /app/booking-service /app/booking-service

CMD ["./booking-service"]