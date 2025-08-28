# استفاده از OpenTelemetry Collector Contrib (شامل پلاگین‌های اضافی)
FROM otel/opentelemetry-collector-contrib:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom OpenTelemetry Collector with predefined config"

# کپی فایل کانفیگ
COPY ./config/otel-collector-config.yaml /etc/otel-collector-config.yaml

# فرمان اجرا
CMD ["--config=/etc/otel-collector-config.yaml"]
