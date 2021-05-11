FROM alpine:3.6

ENV APPPATH /opt/app
RUN mkdir -p $APPPATH

COPY ./app-linux $APPPATH/app
RUN chmod +x $APPPATH/app

WORKDIR $APPPATH

EXPOSE 8001

CMD ["./app"]

