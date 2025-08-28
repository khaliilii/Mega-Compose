FROM mysql:latest

LABEL maintainer="Mohsen Khalili"
LABEL description="Custom Postgresql with user and database"

ENV POSTGRES_PASSWORD=mysecretpassword
ENV PGDATA=/var/lib/postgresql/data/pgdata

WORKDIR /var/lib/mysql
RUN chown -R mysql:mysql /var/lib/mysql

EXPOSE 5432

VOLUME ["/var/lib/postgresql/data"]

CMD ["postgres"]
