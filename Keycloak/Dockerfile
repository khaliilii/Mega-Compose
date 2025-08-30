# Base image: Keycloak 21.1.1
FROM quay.io/keycloak/keycloak:21.1.1

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Keycloak IAM Server"

# Volume for persistent data storage
VOLUME ["/opt/keycloak/data"]

# Default command (development mode)
CMD ["start-dev"]
