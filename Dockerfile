FROM golang:1.10
ADD . /go/src/demo-svc
RUN go get -d github.com/gorilla/mux
RUN go install demo-svc

FROM alpine:latest
COPY --from=0 /go/bin/demo-svc .
ENV PORT 8080
CMD ["./demo-svc"]