# Build the service based on the official golang docker image
FROM golang:1.4.2

# Copy the local package files to the container's workspace
ADD . /go/src/

# Build the services
RUN go install consul.demo/fep

# Run the service when container starts
ENTRYPOINT /go/bin/fep

# Expose port 32768 and 5000
EXPOSE 32768 5000
