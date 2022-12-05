# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:latest as builder

# Add Maintainer Info
LABEL maintainer="Md. Ishtiaq Islam <islam.ishtiaq99@gmail.com>"

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source from the current directory to the working directory inside the container
Copy . .

# Build the GO app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o apiserver .

######## Start a new stage from scratch #######
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/apiserver .

# Expose port 8080 to the outside world
EXPOSE 3000

# Command to run the executable
CMD ["./apiserver", "start"]