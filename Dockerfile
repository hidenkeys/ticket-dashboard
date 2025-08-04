# Step 1: Build the Go binary
FROM golang:1.23 AS build

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./

# Install Go dependencies
RUN go mod tidy

# Copy the rest of the application source code
COPY . .

# Install dependencies from the Makefile (like oapi-codegen)
RUN make install

# Generate the API code (if necessary, based on the Makefile)
RUN make api-generate

# Build the Go application (this assumes the binary is named `ticket-dashboard`)
RUN GOOS=linux GOARCH=amd64 go build -o ticket-dashboard main.go

# Step 2: Create a smaller image to run the Go application
FROM alpine:latest

# Install required certificates (needed for HTTPS requests, if applicable)
RUN apk --no-cache add ca-certificates

# Set the Working Directory inside the container
WORKDIR /root/

# Copy the Go binary from the build image
COPY --from=build /app/ticket-dashboard .

COPY .env .env

# Expose the port that the application will run on (adjust port if necessary)
EXPOSE 8082

# Start the Go application
CMD ["./ticket-dashboard"]
