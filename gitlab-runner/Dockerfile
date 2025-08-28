# Dockerfile برای GitLab Runner
FROM gitlab/gitlab-runner:latest

# اضافه کردن دستورات سفارشی (اختیاری)
# مثلاً register خودکار runner با GitLab
# RUN gitlab-runner register --non-interactive \
#   --url "http://gitlab.local" \
#   --registration-token "TOKEN" \
#   --executor "docker" \
#   --docker-image "docker:20.10.16" \
#   --description "docker-runner" \
#   --tag-list "docker,ci" \
#   --run-untagged="true" \
#   --locked="false"

# expose پورت اختیاری
EXPOSE 8093
