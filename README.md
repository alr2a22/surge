# Surge

This project is part of Snapp interview. The Task description is available [here](https://github.com/AliKarami/interview-tasks/tree/master/surge).

## Features:

- Calculate rate coeffiect of a ride based on demand rate in a district
- Thresholds/Coefficients should be configurable

Also for a PoC I consider these features:

- Authentication
- Admin level access for configurable apis

## Notes (my opinions on project)

- This is a real time service and time is so valuable for us
- To achive the district of requested location we need geographic data. API call is an expensive action. Can be solved like a Geometry problem: Find which polygon contains point among a set of polygons.
- RideRequest data for storing a log like, timeseries and have spital data
  - log: large and without update or delete
  - timeseries: query on duration and have fresh data using TSDBs is so useful
  - spital: query on a area for analytic or other queries that work with surge project(like get ride requests near a driver)
- Demand aggrigator for a district can be temporary achived in a cache and can be like a rate limiter algorithem (like Sliding Window rate limiter). Using [Sliding Window](https://www.geeksforgeeks.org/window-sliding-technique/) have some cons and pros:

  - pros: super fast query time with redis
  - cons: result have low accuracy compare to direct query from database. The smaller window size decreases error.

  _for me this(using cache) is a speed-accuracy tradeoff and using TSDBs(like cassandra) maybe is a better approach_

- This service only have destination so real rate for journey should come from other services. Athentication can be in a separate service.

## Design

- Postgis: PostGIS is a spatial database extender for PostgreSQL object-relational database. It adds support for geographic objects allowing location queries to be run in SQL. At first i choose MariaDB. Because of documentation and community was switched to PostGIS.
- Redis: Redis is an in-memory data structure store, used as a distributed, in-memory key–value database, cache and message broker, with optional durability.
- Adminer: Adminer is a tool for managing content in databases.
- query to calculate demand with PostGIS in [here](./api/ride/models/ride_req.go) and with redis in [here](./internal/db/redis.go).
- OpenStreetMap: To get district polygons and import it to PostGIS run [query.txt](./data/query.txt) on [overpass](https://overpass-turbo.eu/#) site and export geojson data then use [import.sh](./scripts/import.sh) with proper database config.

![components](docs/components.jpg?raw=true)

## Usage

For production follow these steps:

1. First fill envT file with correct values and create .env file next to it. for example:

   ```bash
   REDIS_HOST=cache-redis
   REDIS_PASSWORD=1234
   REDIS_PORT=6379
   POSTGRES_USER=user
   POSTGRES_PASSWORD=1234
   POSTGRES_DB=db
   POSTGRES_HOST=db-postgres
   POSTGRES_PORT=5432
   LOGLEVEL=info
   JWT_SECRET=asdf1234
   JWT_VALID_DAYS=10
   REGISTRY=registry.docker.ir
   WINDOW_MINUTES=10
   ```

2. The command below run all commponents of project
   ```bash
   make prod-compose-up
   ```
3. To create admin user run
   ```bash
   make prod-admin username=admin password=admin
   ```
4. To import geojson file run

   ```bash
   make import-geojson
   ```

5. To add default thresholds
   ```bash
   make add-default-thresholds
   ```

Also for development all commands available in [Makefile](./Makefile)

## Postman

All REST API endpoints available in docs directory [here](./docs/surge.postman_collection.json).
