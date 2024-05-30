#
# First stage:
# Building a frontend.
#

FROM alpine:3.17 AS frontend

# Move to a working directory (/static).
WORKDIR /static

# https://stackoverflow.com/questions/69692842/error-message-error0308010cdigital-envelope-routinesunsupported
ENV NODE_OPTIONS=--openssl-legacy-provider
# Install npm (with latest nodejs)
RUN apk add --update nodejs npm && \
    npm i -g -s --unsafe-perm

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

# Copy and download dependencies.
COPY go.mod go.sum ./
RUN go mod download

# Copy a source code to the container.
COPY . .

# Copy frontend static files from /static to the root folder of the backend container.
COPY --from=frontend ["/static/dist", "ui/dist"]

RUN make GCFLAGS="-tags=prod"

#
# Third stage:
# Creating and running a new scratch container with the backend binary.
#

FROM scratch

# Copy binary from /build to the root folder of the scratch container.
COPY --from=backend ["/build/fitwave", "/"]

# Command to run when starting the container.
ENTRYPOINT ["/fitwave"]
