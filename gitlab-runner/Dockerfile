# Base image: GitLab Runner
FROM gitlab/gitlab-runner:latest

# Optional: add custom commands
# For example, auto-register runner with GitLab
# RUN gitlab-runner register --non-interactive \
#   --url "http://gitlab.local" \
#   --registration-token "TOKEN" \
#   --executor "docker" \
#   --docker-image "docker:20.10.16" \
#   --description "docker-runner" \
#   --tag-list "docker,ci" \
#   --run-untagged="true" \
#   --locked="false"

# Optional: expose port
EXPOSE 8093
