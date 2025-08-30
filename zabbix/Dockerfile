
FROM zabbix/zabbix-server-mysql:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Zabbix Server for monitoring"

# Volume Mysql Check
VOLUME ["/var/lib/mysql"]

CMD ["zabbix_server", "-f"]
