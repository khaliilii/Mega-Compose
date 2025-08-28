# Swagger UI Dockerfile
FROM swaggerapi/swagger-ui:v5.11.0

# Copy your OpenAPI spec (swagger.json or openapi.yaml)
COPY swagger.json /usr/share/nginx/html/swagger.json

# Configure Swagger UI to load your spec by default
ENV SWAGGER_JSON=/usr/share/nginx/html/swagger.json


# # Build image
# docker compose build swagger

# # Start service
# docker compose up -d swagger
