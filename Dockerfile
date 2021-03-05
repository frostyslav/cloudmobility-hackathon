# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16.0-buster as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/frostyslav/cloudmobility-hackathon
COPY . .

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN cd cmd/knfunc && CGO_ENABLED=0 GOOS=linux go build -v -a -o /knfunc

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.13.2
# RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /knfunc /knfunc

# Run the web service on container startup.
CMD ["/knfunc"]
