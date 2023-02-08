FROM golang:1.19-alpine3.16 as builder
RUN apk add --no-cache git
RUN apk add --no-cache sqlite-libs sqlite-dev
RUN apk add --no-cache build-base
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get -d -v
RUN go build -v -o main .

FROM alpine:3.17
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/main .
COPY --from=builder /app/config.json .

CMD ["/root/main"]
