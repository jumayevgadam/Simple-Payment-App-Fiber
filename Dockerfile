FROM golang:1.23.3

WORKDIR /app
 
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /tsu-toleg

EXPOSE 8080

CMD [ "tsu-toleg" ]