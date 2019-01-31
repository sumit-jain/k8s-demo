#FROM golang:1.10 AS build-env
#ADD . /go/src/demo-svc
#RUN go get -d github.com/gorilla/mux
#RUN go install demo-svc
#FROM alpine
#WORKDIR /go/src/demo-svc
#COPY --from=build-env /go/src/demo-svc /go/src/demo-svc
#RUN chown nobody:nogroup /go
#USER nobody
#ENTRYPOINT /go/bin/demo-svc
#ENV PORT 8080

############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
# Fetch dependencies.
# Using go get.
RUN go get -d -v
# Build the binary.
RUN go build -o /go/bin/hello
############################
# STEP 2 build a small image
############################
FROM scratch
# Copy our static executable.
COPY --from=builder /go/bin/hello /go/bin/hello
# Run the hello binary.
ENTRYPOINT ["/go/bin/hello"]
