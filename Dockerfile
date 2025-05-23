# Use ubuntu as the base image
FROM docker.io/library/ubuntu:latest

# Install necessary packages: Go, Git, etc.
RUN apt-get update && \
    apt-get install -y golang-go git && \
    apt-get clean

# Set the working directory inside the container
WORKDIR /app

# Copy the Go source code into the container
COPY main.go .

# Copy the Go source code into the container
COPY rajni_ascii .

# Build the Go application
RUN go build -o app main.go

# Run the binary by default
ENTRYPOINT ["./app"]