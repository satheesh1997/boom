FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /go/src/pkg.satheesh.dev/boom

COPY . .

# Environment variables required for our image to build.
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor

# Fetch dependencies.
# Using go mod with go 1.11.
RUN go mod download
RUN go mod vendor
RUN go mod verify

# Build the binary.
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/boom

##############################
#     Alpine 3.16 Image      #
##############################

FROM alpine:3.16

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/bin/boom /main

# Environment variables required for our image to run.
ENV GIN_MODE=release

ENTRYPOINT ["/main"]
