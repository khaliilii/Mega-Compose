FROM nginx:1.9

LABEL MAINTAINER="Mohsen khalili <khalili.mohsen.mk@gmail.com>"

# حذف محتوای پیش‌فرض nginx
RUN rm -rf /usr/share/nginx/html/*

# کپی فایل‌های تنظیمات
COPY configs/nginx.conf /etc/nginx/nginx.conf
COPY configs/default.conf /etc/nginx/conf.d/default.conf

# کپی گواهینامه‌ها و کلیدها
COPY certs/productionexample.crt /etc/ssl/certs/productionexample.crt
COPY certs/productionexample.key /etc/ssl/private/productionexample.key
COPY certs/dhparam.pem /etc/ssl/private/dhparam.pem

# کپی اسکریپت entrypoint
COPY docker-entrypoint /docker-entrypoint
RUN chmod +x /docker-entrypoint

# ست کردن entrypoint و فرمان پیش‌فرض
ENTRYPOINT ["/docker-entrypoint"]
CMD ["nginx", "-g", "daemon off;"]
