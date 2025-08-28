# 🚀 Mega Compose – Awesome Docker Services

An **all-in-one DevOps & Cloud-Native Lab** built with **Docker Compose**.  
This repository includes multiple categories of services for **databases, monitoring, storage, CI/CD, collaboration, reverse proxy** and more.  
Perfect for testing, learning, and building a complete **DevOps Lab** on Docker.

---

## 📦 Services

### 🗄️ Databases
- MySQL  
- PostgreSQL  
- MongoDB  
- Redis / Valkey  
- Memcached  

### 📚 Database GUI Tools
- phpMyAdmin (MySQL GUI)  
- pgAdmin (PostgreSQL GUI)  
- Mongo Express (MongoDB GUI)  
- RedisInsight (Redis GUI)  

### 🐇 Messaging & Streaming
- RabbitMQ  
- Kafka + Zookeeper  

### ⚙️ Application Runtimes
- Backend Core (backApp)  
- Frontend App (frontApp)  
- Go App (goApp)  
- php-nginx  
- php-apache  

### 📊 Monitoring & Observability
- Prometheus  
- Grafana  
- cAdvisor  
- Jaeger  
- OpenTelemetry Collector (OTEL)  
- Zabbix  

### 🔍 Logging & Search
- Elasticsearch  
- Logstash  
- Kibana  
- Filebeat  

### 🛠️ DevOps Tools
- Portainer  
- Rancher  
- SonarQube (Code Quality)  
- Restic (Backup)  
- Kopia (Backup)  
- Trivy (Security Scanner)  

### 🗂️ Storage & File Sharing
- Nextcloud  
- MinIO  
- Ceph  
- Longhorn  

### 🔑 Identity & Access
- Keycloak  

### 📬 Mail & Messaging
- Mailpit (SMTP testing)  
- Rocket.Chat (Team Chat)  

### 🦾 CI/CD & Source Control
- GitLab CE  
- GitLab Runner  
- Gitea  

### 🌐 CMS & Collaboration
- WordPress  
- OpenProject  
- Wiki.js  

### 🔒 Networking & Proxy
- Traefik  
- Nginx  

---

## ⚡ How to Use

### ▶️ Run All Services (Mega Compose)
To spin up **all services at once**:
```bash
docker compose up -d
