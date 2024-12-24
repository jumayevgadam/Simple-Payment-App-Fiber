FROM golang:1.23-bookworm

# Set the working directory inside the container
WORKDIR /app

COPY . .

RUN go mod tidy

WORKDIR /app

RUN go build cmd/main.go

EXPOSE 7000

CMD ["go","run","cmd/main.go"]