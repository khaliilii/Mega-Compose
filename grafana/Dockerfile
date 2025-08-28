# استفاده از ایمیج رسمی Grafana
FROM grafana/grafana:latest

# اگر فایل‌های پیش‌فرض داشبورد یا provisioning داری
# COPY provisioning /etc/grafana/provisioning
# COPY dashboards /var/lib/grafana/dashboards

# داده‌های Grafana روی volume
VOLUME ["/var/lib/grafana"]

EXPOSE 3000

# دستور پیش‌فرض Grafana
CMD ["/run.sh"]
