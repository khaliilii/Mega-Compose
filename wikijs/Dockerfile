# Base image: Wiki.js official
FROM requarks/wiki:2

# Set working directory
WORKDIR /wiki

# Copy custom configuration (if you have one)
# Custom config file copy
COPY config.yml /wiki/config.yml

# Copy custom branding / themes if needed
# COPY ./themes /wiki/themes

# Expose default Wiki.js port
EXPOSE 3000

# Start Wiki.js
CMD ["node", "server"]
