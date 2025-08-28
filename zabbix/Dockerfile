# Dockerfile برای Zabbix Server
FROM zabbix/zabbix-server-mysql:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Zabbix Server for monitoring"

# Volume برای دیتابیس و داده‌ها
VOLUME ["/var/lib/mysql"]

CMD ["zabbix_server", "-f"]
