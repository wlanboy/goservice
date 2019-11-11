FROM golang:1.13.4

WORKDIR /app
COPY .env .env
COPY goservice goservice
EXPOSE 8000

CMD ["./goservice"]