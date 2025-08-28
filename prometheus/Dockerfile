# استفاده از ایمیج رسمی
FROM prom/prometheus:latest

# کپی فایل کانفیگ Prometheus به داخل کانتینر
COPY prometheus.yml /etc/prometheus/prometheus.yml

# داده‌های Prometheus روی volume
VOLUME ["/prometheus"]

EXPOSE 9090

# دستور اصلی Prometheus
CMD [ "--config.file=/etc/prometheus/prometheus.yml", "--storage.tsdb.path=/prometheus" ]
