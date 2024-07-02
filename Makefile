DOCKER_COMPOSE=docker_compose.yaml
ENV=.env

.PHONY: run build up down ps version logs stop

run: build up logs

build:
	docker-compose --file $(DOCKER_COMPOSE) --env-file $(ENV) build

up:
	docker-compose --file $(DOCKER_COMPOSE) --env-file $(ENV) up -d

logs:
	docker-compose --file $(DOCKER_COMPOSE) logs -t

#config:
#	docker-compose --file $(DOCKER_COMPOSE) --env-file $(ENV) config

down:
	docker-compose --file $(DOCKER_COMPOSE) --env-file $(ENV) down

ps:
	docker-compose --file $(DOCKER_COMPOSE) ps

version:
	docker-compose version

stop:
	docker-compose --file $(DOCKER_COMPOSE) stop