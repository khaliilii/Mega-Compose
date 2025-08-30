# Base: Alpine for small footprint
FROM alpine:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom Kopia backup tool container"

# Install required packages
RUN apk add --no-cache curl tar bash

# Download Kopia CLI
ENV KOPIA_VERSION=0.14.3
RUN curl -L https://github.com/kopia/kopia/releases/download/v${KOPIA_VERSION}/kopia-${KOPIA_VERSION}-linux-x64.tar.gz \
    | tar -xz -C /usr/local/bin kopia

# Create backup storage folder
RUN mkdir -p /data

# Expose REST API port if needed
EXPOSE 51515

# Set working directory
WORKDIR /data

# Default command (start REST server)
CMD ["kopia", "server", "start", "--address", "0.0.0.0:51515", "--base-directory", "/data"]
