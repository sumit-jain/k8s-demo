FROM golang:1.10 AS build-env
COPY . /app
RUN go get -d github.com/gorilla/mux
RUN cd /app && go build -o demo-svc main.go
FROM alpine
WORKDIR /app
COPY --from=build-env /app/demo-svc /app/demo-svc
RUN chown nobody:nogroup /app
USER nobody
ENTRYPOINT /app/demo-svc
ENV PORT 8080
