# Base image: Jaeger all-in-one 1.46
FROM jaegertracing/all-in-one:1.46

LABEL maintainer="Your Name <your@example.com>"
LABEL description="Jaeger All-in-One for distributed tracing"

# Optional volume for persistent data
VOLUME ["/jaeger"]

# Default command (expose Zipkin HTTP port 9411)
CMD ["--collector.zipkin.http-port=9411"]
