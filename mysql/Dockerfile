FROM mysql:latest

LABEL maintainer="Mohsen Khalili"
LABEL description="Custom MySQL with user and database"

ENV MYSQL_ROOT_PASSWORD=rootpass \
    MYSQL_DATABASE=my_database \
    MYSQL_USER=external_user \
    MYSQL_PASSWORD=password

WORKDIR /var/lib/mysql
RUN chown -R mysql:mysql /var/lib/mysql

RUN mysql -u root -p -e "CREATE USER '$MYSQL_USER'@'%' IDENTIFIED BY '$MYSQL_PASSWORD';"
RUN mysql -u root -p -e "GRANT ALL PRIVILEGES ON $MYSQL_DATABASE.* TO '$MYSQL_USER'@'%';"
RUN mysql -u root -p -e "FLUSH PRIVILEGES;"

EXPOSE 3306

VOLUME ["/var/lib/mysql"]

CMD ["mysqld"]
