FROM golang:alpine AS builder
RUN apk update && apk upgrade && \
    apk add curl wget
RUN mkdir -p /app
WORKDIR /app
EXPOSE 8000
ADD simple_webserver.go .
RUN go build -o webserver .

FROM alpine
WORKDIR /app
COPY --from=builder /app/ /app/
CMD ["./webserver"]
