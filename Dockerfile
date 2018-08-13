FROM golang:1.8-alpine
ADD . /go/src/demo-svc
RUN go get github.com/gorilla/mux
RUN go install demo-svc

FROM alpine:latest
COPY --from=0 /go/bin/demo-svc .
ENV PORT 8080
CMD ["./demo-svc"]