# Use official OpenProject Community Edition image as base
FROM openproject/community:latest

# Set maintainer
LABEL maintainer="yourname@example.com"

# Optional environment variables
ENV SECRET_KEY_BASE=your-random-secret
ENV DATABASE_URL=postgres://openproject:password@db/openproject

# Default OpenProject port
EXPOSE 8080

# Default command to run OpenProject
CMD ["openproject", "run"]
