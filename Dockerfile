FROM alpine

WORKDIR /srv/investool

ADD ./dist/investool.tar.gz /srv/

EXPOSE 9527 4870
ENTRYPOINT ["./investool", "-c", "./config.toml"]
