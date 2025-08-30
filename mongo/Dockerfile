# Base image
FROM mongo:6

LABEL maintainer="Your Name <your@email.com>"

# Set environment variables for Mongo root user
ENV MONGO_INITDB_ROOT_USERNAME=root \
    MONGO_INITDB_ROOT_PASSWORD=rootpassword \
    MONGO_INITDB_DATABASE=my_database

# Expose default MongoDB port
EXPOSE 27017

# Command to run MongoDB
CMD ["mongod"]
