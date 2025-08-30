FROM axllent/mailpit:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom Mailpit container for local email testing"

# Expose ports:
# 1025 -> SMTP
# 8025 -> Web UI
EXPOSE 1025 8025

# Default command
CMD ["mailpit", "--smtp", "0.0.0.0:1025", "--ui", "0.0.0.0:8025"]
