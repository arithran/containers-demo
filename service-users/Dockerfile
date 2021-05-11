FROM alpine:3.6

ENV SERVICE_PASSWORD_ADDRESS=http://service-passwords:8001

ENV APPPATH /opt/app
RUN mkdir -p $APPPATH

COPY ./app-linux $APPPATH/app
RUN chmod +x $APPPATH/app

WORKDIR $APPPATH

EXPOSE 8000

CMD ["./app"]

