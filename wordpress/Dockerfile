# image load
FROM wordpress:php8.2-apache

# Add env
ENV WORDPRESS_DB_HOST=db:3306
ENV WORDPRESS_DB_USER=wordpress
ENV WORDPRESS_DB_PASSWORD=wordpresspass
ENV WORDPRESS_DB_NAME=wordpressdb

# ports
EXPOSE 80

# default command
CMD ["apache2-foreground"]
