FROM busybox:1.31

RUN mkdir /home
COPY ./goservice /home/
COPY ./.env /home/
EXPOSE 8000

CMD ["/home/goservice"]
