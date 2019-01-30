FROM golang:1.10 AS build-env
ADD . /go/src/demo-svc
RUN go get -d github.com/gorilla/mux
RUN go install demo-svc
FROM alpine
WORKDIR /app
COPY --from=build-env /go/src/demo-svc /go/src/demo-svc
RUN chown nobody:nogroup /app
USER nobody
ENTRYPOINT /go/bin/demo-svc
ENV PORT 8080
