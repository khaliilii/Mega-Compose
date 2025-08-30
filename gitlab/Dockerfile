# Base image: GitLab Community Edition
FROM gitlab/gitlab-ce:latest

# Optional additional configuration
ENV GITLAB_OMNIBUS_CONFIG="external_url 'http://gitlab.local'; gitlab_rails['gitlab_shell_ssh_port']=2224"

# Expose GitLab ports
EXPOSE 80 443 2224

# Volumes for persistent data
VOLUME ["/etc/gitlab", "/var/opt/gitlab", "/var/log/gitlab"]
