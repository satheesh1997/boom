FROM golang:alpine AS builder

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

WORKDIR /go/src/github.com/satheesh1997/boom

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
# STEP 2 build a small image #
##############################

FROM alpine:latest

RUN apk --no-cache add ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /go/bin/boom /boom

# Copy the template files to the production image from the builder stage.
COPY --from=builder /go/src/github.com/satheesh1997/boom/templates /templates

# Environment variables required for our image to run.
ENV GIN_MODE=release

# Run the binary.
ENTRYPOINT ["/boom"]
