# Build stage
FROM golang:1.23 AS builder

WORKDIR /app

COPY ./ ./

ENV GO111MODULE=on

RUN go build cmd/main.go

# Final stage
FROM alpine:3.17.2

WORKDIR /app

COPY --from=builder /app ./

EXPOSE 4000

CMD [ "./main" ]
