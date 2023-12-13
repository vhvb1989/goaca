# syntax=docker/dockerfile:1
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /app

# Download Go modules (dependencies)
COPY go.mod ./
#RUN go mod download

# Copy the source code into the container
COPY . .

# Expose the port on which the server will listen
EXPOSE 3100

# Build and run the Go application
CMD ["go", "run", "main.go"]
