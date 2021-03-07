# Use a minimal base image with Go support.
FROM golang:alpine AS build

# Install cURL so that we can use it to define health checks.
RUN ["apk", "add", "curl"]

# Copy in source code and build the server.
COPY ["describer.go", "/describer.go"]
RUN ["go", "build", "-o", "/describer", "/describer.go"]

# Set the entrypoint.
ENTRYPOINT ["/describer"]
