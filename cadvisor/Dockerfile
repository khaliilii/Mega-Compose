# Base image: official cAdvisor
FROM gcr.io/cadvisor/cadvisor:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="cAdvisor for monitoring Docker containers"

# Expose default cAdvisor port
EXPOSE 8080

# Default command
CMD ["/usr/bin/cadvisor", "-port=8080", "-docker_only"]
