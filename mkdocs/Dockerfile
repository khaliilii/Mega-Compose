# Dockerfile
FROM python:3.10-slim

# جلوگیری از pyc و بهتر شدن لاگ‌ها
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

# نصب وابستگی‌ها
RUN apt-get update && apt-get install -y \
    build-essential \
    git \
    && rm -rf /var/lib/apt/lists/*

# نصب MkDocs + تم محبوب Material
RUN pip install --no-cache-dir \
    mkdocs \
    mkdocs-material \
    mkdocs-mermaid2-plugin

# فولدر پروژه
WORKDIR /docs

# پورت MkDocs
EXPOSE 8000

# اجرای MkDocs به صورت live server
CMD ["mkdocs", "serve", "-a", "0.0.0.0:8000"]


# docker compose up -d
# http://localhost:8000
