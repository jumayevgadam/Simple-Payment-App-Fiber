FROM golang:1.23.3

# Set the working directory inside the container
WORKDIR /tsu-toleg

# Copy the entire project into the container
COPY . .

# Explicitly enable Go modules
ENV GO111MODULE=on

# Build the go application
RUN go build -o main cmd/main.go

# EXPOSE the port the application runs on
EXPOSE 8080

# Command to run go application
CMD [ "tsu-toleg/main" ]