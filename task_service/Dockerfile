# Use the official Golang image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the application source code
COPY . .

# Build the Go application
RUN go build -o server .

# Expose the port on which your application runs
EXPOSE 9000

# Define the command to run your application
CMD ["./server"]