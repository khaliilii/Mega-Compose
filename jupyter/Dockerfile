# Dockerfile
FROM python:3.10-slim

# جلوگیری از ایجاد pyc
ENV PYTHONDONTWRITEBYTECODE=1
ENV PYTHONUNBUFFERED=1

# نصب وابستگی‌ها
RUN apt-get update && apt-get install -y \
    build-essential \
    curl \
    && rm -rf /var/lib/apt/lists/*

# نصب Jupyter و کتابخانه‌ها
RUN pip install --no-cache-dir \
    jupyter \
    numpy \
    pandas \
    matplotlib \
    scikit-learn

# ساخت پوشه کاری
WORKDIR /workspace

# پورت Jupyter
EXPOSE 8888

# اجرای Jupyter Notebook
CMD ["jupyter", "notebook", "--ip=0.0.0.0", "--port=8888", "--no-browser", "--allow-root"]
