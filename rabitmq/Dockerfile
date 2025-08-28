# Dockerfile برای RabbitMQ با Management UI
FROM rabbitmq:3-management

# نصب پلاگین‌های اضافی (اختیاری)
RUN rabbitmq-plugins enable --offline rabbitmq_federation rabbitmq_federation_management

# کپی فایل کانفیگ سفارشی (اختیاری)
# COPY rabbitmq.conf /etc/rabbitmq/rabbitmq.conf

# پورت‌ها (برای reference، container خودش expose می‌کنه)
EXPOSE 5672 15672
