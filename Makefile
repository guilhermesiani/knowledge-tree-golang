.PHONY: up
up:
	docker-compose up -d --build

.PHONY: down
down:
	docker-compose down

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: test
test:
	docker-compose exec app go test ./...

.PHONY: deep-test
deep-test:
	curl http://127.0.0.1:3001/myself -v