# Use the Go official image
# https://hub.docker.com/_/golang
FROM golang:1.22
 
# Create and change to the app directory.
WORKDIR /app
 
# Copy local code to the container image.
COPY . ./

# Install project dependencies
RUN go mod download

# Build the app
RUN go build -o app

# Run the service on container startup.
ENTRYPOINT ["./app"]