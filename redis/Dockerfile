FROM redis:latest

LABEL maintainer="Mahmoud Zalt <mahmoud@zalt.me>"

## For security settings uncomment, make the dir, copy conf, and also start with the conf, to use it
RUN mkdir -p /etc/redis
COPY redis.conf /etc/redis/redis.conf

VOLUME /data

CMD ["redis-server"]
#CMD ["redis-server"]

EXPOSE 6379