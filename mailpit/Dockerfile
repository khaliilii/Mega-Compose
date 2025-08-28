# استفاده از آلباین برای سبک بودن
FROM alpine:3.20

LABEL maintainer="Your Name <your@email.com>"
LABEL description="Mailpit Docker Image"

# نصب ابزارهای ضروری
RUN apk add --no-cache curl ca-certificates tzdata

# دانلود آخرین نسخه Mailpit
ENV MAILPIT_VERSION=v1.22.1

RUN curl -L -o /usr/local/bin/mailpit \
    https://github.com/axllent/mailpit/releases/download/${MAILPIT_VERSION}/mailpit-linux-amd64 \
    && chmod +x /usr/local/bin/mailpit

# پورت‌های مورد نیاز
EXPOSE 1025 8025

# اجرای سرویس
ENTRYPOINT ["mailpit"]
CMD ["--listen=0.0.0.0:1025", "--ui-bind=0.0.0.0:8025"]
