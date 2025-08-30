FROM nginx:1.9

LABEL MAINTAINER="Mohsen Khalili <khalili.mohsen.mk@gmail.com>"

# Remove default Nginx content
RUN rm -rf /usr/share/nginx/html/*

# Copy configuration files
COPY configs/nginx.conf /etc/nginx/nginx.conf
COPY configs/default.conf /etc/nginx/conf.d/default.conf

# Copy certificates and keys
COPY certs/productionexample.crt /etc/ssl/certs/productionexample.crt
COPY certs/productionexample.key /etc/ssl/private/productionexample.key
COPY certs/dhparam.pem /etc/ssl/private/dhparam.pem

# Copy entrypoint script
COPY docker-entrypoint /docker-entrypoint
RUN chmod +x /docker-entrypoint

# Set entrypoint and default command
ENTRYPOINT ["/docker-entrypoint"]
CMD ["nginx", "-g", "daemon off;"]
