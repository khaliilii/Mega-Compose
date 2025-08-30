# Use OpenTelemetry Collector Contrib (includes extra plugins)
FROM otel/opentelemetry-collector-contrib:latest

# Maintainer info
LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom OpenTelemetry Collector with predefined config"

# Copy custom configuration file
COPY ./config/otel-collector-config.yaml /etc/otel-collector-config.yaml

# Command to run the collector
CMD ["--config=/etc/otel-collector-config.yaml"]
