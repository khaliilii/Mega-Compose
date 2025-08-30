# Base image: official Kafka 3.5.1
FROM wurstmeister/kafka:3.5.1

LABEL maintainer="Mohsen Khalili <khalili.mohsen.mk@gmail.com>"
LABEL description="Custom Kafka image with Debezium connectors"

# Switch to root to install additional tools
USER root

# Create folder for Kafka connectors
RUN mkdir -p /kafka/connectors

# Example: add Debezium MySQL connector
# Copy your connector JAR file
COPY connectors/debezium-connector-mysql-1.9.7.jar /kafka/connectors/

# Set environment variable for Kafka Connect plugin path
ENV KAFKA_CONNECT_PLUGIN_PATH=/kafka/connectors

# Set correct permissions for connector folder
RUN chown -R kafka:kafka /kafka/connectors

# Switch back to Kafka user
USER kafka

# Default command from base image
CMD ["start-kafka.sh"]
