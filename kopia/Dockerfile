FROM alpine:latest

# آخرین نسخه Kopia (می‌تونی ورژن دقیق هم بذاری)
ENV KOPIA_VERSION=0.17.0

RUN apk add --no-cache bash curl ca-certificates fuse3 \
    && curl -L https://github.com/kopia/kopia/releases/download/v${KOPIA_VERSION}/kopia-${KOPIA_VERSION}-linux-x64.tar.gz \
    | tar -xz -C /usr/local/bin \
    && chmod +x /usr/local/bin/kopia*

WORKDIR /data

ENTRYPOINT ["kopia"]

#docker build -t my-kopia .

# docker run --rm \
#   -v /data:/data:ro \
#   -v /backups:/backups \
#   my-kopia \
#   snapshot create /data
# //////////////////////////////
# webui kopia:
# docker run -d --name kopia-ui \
#   -v ./config:/app/config \
#   -v ./cache:/app/cache \
#   -v ./logs:/app/logs \
#   -v /data:/data:ro \
#   -v ./backups:/backups \
#   -p 51515:51515 \
#   my-kopia \
#   server start --insecure --disable-csrf --address=0.0.0.0:51515

