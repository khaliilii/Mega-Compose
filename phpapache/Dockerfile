# Use official PHP-Apache image
FROM php:8.2-apache

# Maintainer info
LABEL maintainer="yourname@example.com"
LABEL description="PHP + Apache Docker image for your application"

# Enable required PHP extensions
RUN docker-php-ext-install mysqli pdo pdo_mysql && \
    docker-php-ext-enable pdo_mysql

# Optional: enable Apache modules
RUN a2enmod rewrite headers

# Set working directory
WORKDIR /var/www/html

# Copy your application files
COPY ./www/ /var/www/html/

# Set permissions (optional)
RUN chown -R www-data:www-data /var/www/html \
    && chmod -R 755 /var/www/html

# Expose Apache port
EXPOSE 80

# Start Apache in the foreground
CMD ["apache2-foreground"]
