# Start from a Debian image with the latest version of Go installed
FROM golang
# Copy the local package files to the container's workspace.
ADD . /go/src/demo-svc
# fetch or manage dependencies
RUN go get -d github.com/gorilla/mux
# Build the demo-svc inside the container.
RUN go install demo-svc

# Run the demo-svc by default when the container starts.
ENTRYPOINT /go/bin/demo-svc

# Document that the service listens on port 8080.
ENV PORT 8080