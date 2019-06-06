FROM golang:alpine
RUN apk update && apk upgrade && \
    apk add curl wget
RUN mkdir -p /app
WORKDIR /app
EXPOSE 8000
ADD . /app
RUN go build /app/simple_webserver.go
CMD ["./simple_webserver"]
