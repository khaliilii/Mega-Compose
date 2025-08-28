# استفاده از آخرین ورژن Valkey
FROM valkey/valkey:latest

LABEL maintainer="Your Name <you@example.com>"
LABEL description="Custom Valkey Docker Image"

# کپی کانفیگ سفارشی
COPY valkey.conf /usr/local/etc/valkey/valkey.conf

# مسیر داده‌ها
VOLUME ["/data"]

# پورت پیش‌فرض
EXPOSE 6379

# اجرای سرور با فایل کانفیگ
CMD ["valkey-server", "/usr/local/etc/valkey/valkey.conf"]
