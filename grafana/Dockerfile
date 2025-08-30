# Use official Grafana image
FROM grafana/grafana:latest

# Optional: copy default provisioning or dashboards
# COPY provisioning /etc/grafana/provisioning
# COPY dashboards /var/lib/grafana/dashboards

# Persist Grafana data
VOLUME ["/var/lib/grafana"]

# Expose Grafana port
EXPOSE 3000

# Default command
CMD ["/run.sh"]
