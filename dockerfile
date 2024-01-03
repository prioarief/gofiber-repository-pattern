# Use the official Go image as the base image
FROM golang:alpine

# Set the working directory in the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

RUN go mod tidy

# Build the Go application inside the container
RUN go build -o main

# Expose a port if your Go app listens on a specific port
EXPOSE 3000

# Define the command to run your Go application
CMD ["./main"]
