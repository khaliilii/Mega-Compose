FROM redis:latest

LABEL org.opencontainers.image.authors="Mohsen Khalili"
LABEL description="Custom redis with password"

COPY ./conf/redis.conf /usr/local/etc/redis/redis.conf

EXPOSE 6379

VOLUME ["/data"]

CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]