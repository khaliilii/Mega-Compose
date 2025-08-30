
# Use official phpMyAdmin image
FROM phpmyadmin/phpmyadmin:latest

# Maintainer info
LABEL maintainer="Mohsen Khalili"
LABEL description="phpMyAdmin container for MySQL/PostgreSQL management"

# Expose default phpMyAdmin port
EXPOSE 80

# Optional environment variables (can be overridden in docker-compose)
ENV PMA_HOST=db
ENV PMA_PORT=3306
ENV UPLOAD_LIMIT=100M
ENV PMA_ARBITRARY=1

# Default command from base image
# CMD ["php-fpm"] or default entrypoint is already set in base image
