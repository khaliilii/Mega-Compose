# Base image: official Apache Airflow
FROM apache/airflow:2.9.1

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom Apache Airflow image with extra Python dependencies"

# Install extra Python packages
USER root
RUN pip install --no-cache-dir \
      pandas \
      requests \
      apache-airflow-providers-postgres

# Switch back to airflow user
USER airflow

# Default command is defined in the base image (airflow webserver/scheduler)
