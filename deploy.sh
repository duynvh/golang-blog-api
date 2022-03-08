#!/usr/bin/env bash
if [ -f .env ]
then
  export $(cat .env | sed 's/#.*//g' | xargs)
fi

echo "Downloading packages..."
go mod download
echo "Compiling..."
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

echo "Docker building..."
docker build -t ${APP_NAME} -f ./Dockerfile .
echo "Docker saving..."
docker save -o ${APP_NAME}.tar ${APP_NAME}

echo "Deploying..."
scp -o StrictHostKeyChecking=no ./${APP_NAME}.tar ${DEPLOY_CONNECT}:~
#ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} 'bash -s' < ./deploy/stg.sh
ssh -o StrictHostKeyChecking=no ${DEPLOY_CONNECT} \
  APP_NAME=${APP_NAME} \
  DOCKER_NETWORK=${DOCKER_NETWORK} \
  EXPORTED_PORT=${EXPORTED_PORT} \
  PORT=${PORT} \
  VIRTUAL_HOST=${VIRTUAL_HOST} \
  LETSENCRYPT_HOST=${LETSENCRYPT_HOST} \
  LETSENCRYPT_EMAIL=${LETSENCRYPT_EMAIL} \
  DB_USER=${DB_USER} \
  DB_PASSWORD=${DB_PASSWORD} \
  DB_HOST=${DB_HOST} \
  DB_NAME=${DB_NAME} \
  S3_BUCKET_NAME=${S3_BUCKET_NAME} \
  S3_REGION=${S3_REGION} \
  S3_API_KEY=${S3_API_KEY} \
  S3_SECRET_KEY=${S3_SECRET_KEY} \
  S3_DOMAIN=${S3_DOMAIN} \
  SYSTEM_SECRET=${SYSTEM_SECRET} \
  'bash -s' < ./deploy/stg.sh

echo "Cleaning..."
rm -f ./${APP_NAME}.tar
#docker rmi $(docker images -qa -f 'dangling=true')
echo "Done"