# Dockerfile برای GitLab CE
FROM gitlab/gitlab-ce:latest

# تنظیمات اضافی (اختیاری)
ENV GITLAB_OMNIBUS_CONFIG="external_url 'http://gitlab.local'; gitlab_rails['gitlab_shell_ssh_port']=2224"

# expose پورت‌ها
EXPOSE 80 443 2224

# حجم‌ها (volume) برای پایدار بودن داده‌ها
VOLUME ["/etc/gitlab", "/var/opt/gitlab", "/var/log/gitlab"]
