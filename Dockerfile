#
# First stage:
# Building a frontend.
#

FROM alpine:3.17 AS frontend

# Move to a working directory (/static).
WORKDIR /static

# https://stackoverflow.com/questions/69692842/error-message-error0308010cdigital-envelope-routinesunsupported
ENV NODE_OPTIONS=--openssl-legacy-provider

# Update repositories and install npm (with latest nodejs)
RUN apk update && apk add --no-cache nodejs npm

# Log npm and node versions for debugging purposes
RUN node -v && npm -v

# Copy only ./ui folder to the working directory.
COPY ui .

# Run npm scripts (install & build).
RUN npm install && npm run build

#
# Second stage:
# Building a backend.
#

FROM golang:1.18-alpine AS backend

# Move to a working directory (/build).
WORKDIR /build

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code to the container.
COPY . .

# Copy frontend static files from /static to the root folder of the backend container.
COPY --from=frontend /static/dist ui/dist

# Build the Go binary
RUN go build -tags=prod -o /build/fitwave ./cmd/fitwave

# Verify the binary is built
RUN ls -l /build/fitwave

#
# Third stage:
# Creating and running a new container with the backend binary.
#

FROM alpine:3.17

# Install necessary runtime dependencies
RUN apk add --no-cache ca-certificates

# Copy binary from /build to the root folder of the container.
COPY --from=backend /build/fitwave /

# Verify the binary is copied correctly
RUN ls -l /fitwave

# Expose the application port
EXPOSE 9000

# Command to run when starting the container.
ENTRYPOINT ["/fitwave"]
