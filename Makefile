include .env

all:
	docker compose up -d --build

down:
	docker compose down

clean: down
	docker rmi grafana/grafana:latest				\
	sportscheduler-api:latest						\
	migrate/migrate:latest 							\
	postgres:17 									\
	prometheuscommunity/postgres-exporter:latest	\
	prom/prometheus:latest 

vclean:
	 docker volume rm sportscheduler_grafana-data 	\
	 sportscheduler_postgres_data                   \
	 sportscheduler_prometheus-data                 
