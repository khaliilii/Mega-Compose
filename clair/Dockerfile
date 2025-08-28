# Clair Dockerfile
FROM quay.io/projectquay/clair:v4.7.2

# Clair runs with a configuration file
COPY config.yaml /config/config.yaml

ENTRYPOINT ["clair"]
CMD ["-config", "/config/config.yaml"]
