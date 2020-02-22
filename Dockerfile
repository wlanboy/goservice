FROM busybox:1.31

COPY ./goservice /home/
EXPOSE 8000

CMD ["/home/goservice"]
