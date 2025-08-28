# Base image Kafka رسمی
FROM wurstmeister/kafka:3.5.1

LABEL maintainer="Mohsen Khalili <khalili.mohsen.mk@gmail.com>"
LABEL description="Custom Kafka image with Debezium connectors"

# نصب ابزارهای اضافی (مثل Debezium)
USER root
RUN mkdir -p /kafka/connectors

# مثال: اضافه کردن Debezium connector MySQL
# اگر داری فایل .jar connector
COPY connectors/debezium-connector-mysql-1.9.7.jar /kafka/connectors/

# Set environment variable for connector plugin path
ENV KAFKA_CONNECT_PLUGIN_PATH=/kafka/connectors

# اجازه دسترسی به پوشه‌ها
RUN chown -R kafka:kafka /kafka/connectors

USER kafka

# Default command از base image استفاده می‌کنه
CMD ["start-kafka.sh"]
