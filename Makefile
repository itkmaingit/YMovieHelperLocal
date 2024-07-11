.PHONY: init
init:
	docker compose up -d --build --remove-orphans

.PHONY: run
run:
	docker compose up -d --remove-orphans

.PHONY: down
down:
	docker compose down
