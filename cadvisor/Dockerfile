# ایمیج رسمی cAdvisor
FROM gcr.io/google-containers/cadvisor:latest

# حتما volume های مورد نیاز برای خواندن metrics کانتینرها رو mount کن
VOLUME ["/rootfs", "/var/run", "/sys", "/var/lib/docker"]

EXPOSE 8080

CMD ["/usr/bin/cadvisor", "-logtostderr"]
