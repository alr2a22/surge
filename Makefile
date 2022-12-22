#!make
include config/.env
export $(shell sed 's/=.*//' config/.env)

server:
	go run main.go run-server

create-adminadmin:
	go run main.go create-admin --username admin --password admin

migrate:
	go run main.go migrate

import-geojson:
	bash scrips/import.sh

dev-compose-up:
	docker-compose --env-file=config/.env up --scale=app=0 -d

prod-compose-up:
	docker-compose --env-file=config/.env up --build -d
