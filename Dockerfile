FROM busybox:latest

COPY ./goservice /home/
EXPOSE 8000

CMD ["/home/goservice"]
