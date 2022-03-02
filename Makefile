include .env

rundb:
	@docker run -d --name mysql --privileged=true -p 3306:3306 \
	-e MYSQL_ROOT_PASSWORD=${DB_ROOT_PASSWORD} \
	-e MYSQL_USER=${DB_USER} \
	-e MYSQL_PASSWORD=${DB_PASSWORD} \
	-e MYSQL_DATABASE=${DB_NAME} \
	-v d:/codes/docker:/bitnami \
	bitnami/mysql:8.0.15

buildmigrator:
	@docker build -t migrator ./migrator

migrateup:
	@docker run --network host migrator -path="/migrations/" -database "mysql://${DSN}" up

startdb:
	@docker start mysql

start:
	@go run .

.PHONY: rundb startdb migrateup buildmigrator start