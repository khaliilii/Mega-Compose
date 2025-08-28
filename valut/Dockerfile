# Vault Dockerfile
FROM hashicorp/vault:1.15.4

# Create config directory
RUN mkdir -p /vault/config

# Copy config file
COPY config.hcl /vault/config/config.hcl

# Expose Vault ports
EXPOSE 8200

# Run Vault
ENTRYPOINT ["vault", "server", "-config=/vault/config/config.hcl"]


# # Build Vault image
# docker compose build vault

# # Start Vault
# docker compose up -d vault

# # Check logs
# docker compose logs -f vault

# docker exec -it vault_service vault kv put secret/hello value=world
