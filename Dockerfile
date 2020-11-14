FROM busybox:1.31

COPY ./goservice /home/
COPY ./.env /home/
EXPOSE 8000

WORKDIR /home
CMD ["/home/goservice"]
