include .env

all:
	docker compose up -d --build

down:
	docker compose down

clean: down
	docker compose down -v
	docker rmi grafana/grafana:latest				\
	sportscheduler-api:latest						\
	migrate/migrate:latest 							\
	postgres:17 									\
	prometheuscommunity/postgres-exporter:latest	\
	prom/prometheus:latest 
