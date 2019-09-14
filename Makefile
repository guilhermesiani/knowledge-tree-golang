.PHONY: up
up:
	docker-compose up -d --build

.PHONY: down
down:
	docker-compose down

.PHONY: test
test:
	curl http://127.0.0.1:3001/friend