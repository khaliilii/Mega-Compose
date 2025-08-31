# 🚀 Mega Compose – Awesome Docker Services

An **all-in-one DevOps & Cloud-Native Lab** built with **Docker Compose**.
This repository includes multiple categories of services for **databases, monitoring, storage, CI/CD, collaboration, reverse proxy** and more.
Perfect for testing, learning, and building a complete **DevOps Lab** on Docker.

---

## 📦 Services

### 🗄️ Databases

* MySQL
* PostgreSQL
* MongoDB
* Redis / Valkey
* Memcached

### 📚 Database GUI Tools

* phpMyAdmin (MySQL GUI)
* pgAdmin (PostgreSQL GUI)
* Mongo Express (MongoDB GUI)
* RedisInsight (Redis GUI)

### 🐇 Messaging & Streaming

* RabbitMQ
* Kafka + Zookeeper

### ⚙️ Application Runtimes

* Backend Core (backApp)
* Frontend App (frontApp)
* Go App (goApp)
* php-nginx
* php-apache
* jupyter

### 📊 Monitoring & Observability

* Prometheus
* Grafana
* cAdvisor
* Jaeger
* OpenTelemetry Collector (OTEL)
* Zabbix

### 🔍 Logging & Search

* Elasticsearch
* Logstash
* Kibana
* Filebeat

### 🛠️ DevOps Tools

* Portainer
* Rancher
* SonarQube (Code Quality)
* Restic (Backup)
* Kopia (Backup)
* Trivy (Security Scanner)
* Clair (Container Security Scanner)
* Vault (Secrets Management)
* Swagger UI (API Documentation)

### 🗂️ Storage & File Sharing

* Nextcloud
* MinIO
* Ceph
* Longhorn

### 🔑 Identity & Access

* Keycloak

### 📬 Mail & Messaging

* Mailpit (SMTP testing)
* Rocket.Chat (Team Chat)

### 🦾 CI/CD & Source Control

* GitLab CE
* GitLab Runner
* Gitea

### 🌐 CMS & Collaboration

* WordPress
* OpenProject
* Wiki.js
* mkdocs
* airflow

### 🔒 Networking & Proxy

* Traefik
* Nginx

---

## ⚡ How to Use

### ▶️ Run All Services (Mega Compose)

To spin up **all services at once**:

```bash
docker compose up -d
```

### 🖥️ Run a Single Service

Each service has its own folder with a **Dockerfile**, so you can run it individually:

```bash
# Example: Run Swagger UI only
cd swagger
docker build -t swagger-ui .
docker run -d -p 8080:8080 swagger-ui
```

```bash
# Example: Run Vault only
cd vault
docker build -t vault .
docker run -d -p 8200:8200 -e 'VAULT_DEV_ROOT_TOKEN_ID=root' vault
```

```bash
# Example: Run Clair only
cd clair
docker build -t clair .
docker run -d -p 6060:6060 clair
```
