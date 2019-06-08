FROM golang:alpine AS builder
RUN mkdir -p /app
WORKDIR /app
EXPOSE 8000
ADD simple_webserver.go .
RUN go build -o webserver .

FROM alpine
RUN apk update && apk upgrade && \
    apk add curl wget
WORKDIR /app
COPY --from=builder /app/ /app/
CMD ["./webserver"]
