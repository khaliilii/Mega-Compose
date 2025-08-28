# پایه: Alpine خیلی سبک
FROM alpine:latest

# نصب Memcached و ابزار لازم
RUN apk add --no-cache memcached su-exec

# ایجاد یوزر امن
RUN adduser -D -g '' memcached

# باز کردن پورت پیش‌فرض
EXPOSE 11211

# اجرای memcached با یوزر امن
USER memcached

# دستور پیش‌فرض
CMD ["memcached", "-u", "memcached", "-m", "64", "-vv"]

# docker build -t my-memcached .

# docker run -d --name memcached -p 11211:11211 my-memcached
