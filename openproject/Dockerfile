# Use official OpenProject Community Edition image as base
FROM openproject/community:latest

# Set maintainer
LABEL maintainer="yourname@example.com"

# اگر بخوای میتونی ENV اضافه کنی
ENV SECRET_KEY_BASE=your-random-secret
ENV DATABASE_URL=postgres://openproject:password@db/openproject

# پورت پیشفرض OpenProject
EXPOSE 8080

# کانتینر به صورت پیشفرض openproject رو اجرا میکنه
CMD ["openproject", "run"]
