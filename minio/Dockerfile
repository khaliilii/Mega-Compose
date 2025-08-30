FROM minio/minio:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom MinIO container for object storage"

# Expose MinIO ports
EXPOSE 9000
EXPOSE 9001

# Default command (server mode)
CMD ["server", "/data", "--console-address", ":9001"]
