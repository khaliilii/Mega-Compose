# Dockerfile برای Jaeger
FROM jaegertracing/all-in-one:1.46

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Jaeger All-in-One for distributed tracing"

# Volume برای داده‌ها (اختیاری)
VOLUME ["/jaeger"]

# فرمان پیش‌فرض تصویر رسمی
CMD ["--collector.zipkin.http-port=9411"]
