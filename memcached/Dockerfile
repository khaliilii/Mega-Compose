FROM memcached:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom Memcached container for caching"

# Expose default Memcached port
EXPOSE 11211

# Run Memcached with custom options (example: 64MB memory, verbose logging)
CMD ["memcached", "-m", "64", "-vv"]
