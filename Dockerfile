# Use the offical Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.16.0-buster as builder

# Copy local code to the container image.
WORKDIR /go/src/github.com/frostyslav/cloudmobility-hackathon
COPY app app
COPY cmd cmd
COPY templates templates
COPY go.mod go.mod
COPY go.sum go.sum

# Build the command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN cd cmd/knfunc && CGO_ENABLED=0 GOOS=linux go build -v -a -o /knfunc

#WORKDIR /
#RUN git clone https://github.com/knative/client.git
#RUN cd client/ && hack/build.sh -f

# Use a Docker multi-stage build to create a lean production image.
# https://docs.docker.com/develop/develop-images/multistage-build/#use-multi-stage-builds
FROM alpine:3.13.2
# RUN apk add --no-cache ca-certificates

RUN wget https://github.com/knative/client/releases/download/v0.21.0/kn-linux-amd64 -O /bin/kn && chmod +x /bin/kn

# Copy the binary to the production image from the builder stage.
COPY --from=builder /knfunc /knfunc
COPY --from=builder /go/src/github.com/frostyslav/cloudmobility-hackathon/templates /templates

# Run the web service on container startup.
CMD ["/knfunc"]
