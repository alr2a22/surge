#!make
include config/.env
export $(shell sed 's/=.*//' config/.env)

migrate: #dev
	go run main.go migrate

create-adminadmin: # dev
	go run main.go create-admin --username admin --password admin

server: # dev
	go run main.go run-server

dev-compose-up: # dev
	docker-compose --env-file=config/.env up --scale=app=0 -d

prod-compose-up: # prod
	docker-compose --env-file=config/.env up --build -d

prod-admin: # prod
	docker exec app /app/main create-admin --username $(username) --password $(password)

import-geojson: # both
	bash scripts/import.sh

add-default-thresholds: # both
	bash scripts/default-thresholds.sh
