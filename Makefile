include .env
export

export PROJECT_ROOT=${shell pwd}

env-up:
	@docker compose up -d penny-plan-postgres

env-down:
	@docker compose down penny-plan-postgres

env-cleanup:
	@read -p "Очистить все volume файлы окружения? Опасность утери данных. [Y/N]: " ans; \
	if [ "$$ans" = "Y" ]; then \
	  make env-down  && \
	  make env-port-close && \
	  rm -rf ${PROJECT_ROOT}/out/pgdata && \
	  echo "Файлы окружения успешно очищены"; \
	else \
	  echo "Очистка окружения отменена"; \
	fi

env-port-forwarder:
	@docker compose up -d port-forwarder

env-port-close:
	@docker compose down port-forwarder

migrate-create:
	@if [ -z "$(seq)" ]; then \
	    echo "Отсутствует обязательный параметр seq. Пример использования make migrate-create seq=init"; \
	    exit 1; \
	fi; \
	docker compose run --rm penny-plan-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
	    echo "Отсутствует обязательный параметр action."; \
		exit 1; \
	fi; \
	docker compose run --rm penny-plan-migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@penny-plan-postgres:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"

app-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/out/logs && \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run ${PROJECT_ROOT}/cmd/penny-plan/main.go