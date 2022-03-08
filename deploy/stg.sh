#!/usr/bin/env bash
docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

docker run -d \
  --name ${APP_NAME} \
  --network ${DOCKER_NETWORK} \
  -p ${EXPORTED_PORT}:${EXPORTED_PORT} \
  -e VIRTUAL_HOST="${VIRTUAL_HOST}" \
  -e LETSENCRYPT_HOST="${LETSENCRYPT_HOST}" \
  -e LETSENCRYPT_EMAIL="${LETSENCRYPT_EMAIL}" \
  -e SYSTEM_SECRET="${SYSTEM_SECRET}" \
  -e DB_USER="${DB_USER}" \
  -e DB_PASSWORD="${DB_PASSWORD}" \
  -e DB_HOST="${DB_HOST}" \
  -e DB_NAME="${DB_NAME}" \
  -e S3_BUCKET_NAME="${S3_BUCKET_NAME}" \
  -e S3_REGION="${S3_REGION}" \
  -e S3_API_KEY="${S3_API_KEY}" \
  -e S3_SECRET_KEY="${S3_SECRET_KEY}" \
  -e S3_DOMAIN="${S3_DOMAIN}" \
  -e APP_ENV="production" \
  ${APP_NAME}