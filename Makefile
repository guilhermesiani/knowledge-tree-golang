.PHONY: build
build:
	docker-compose build

.PHONY: up
up:
	docker-compose up -d

.PHONY: reset
reset:
	docker-compose down && docker-compose up -d

.PHONY: restart
restart:
	docker-compose restart app

.PHONY: get
get:
	docker-compose exec app go get ./...

.PHONY: down
down:
	docker-compose down

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: bash
bash:
	docker-compose exec app bash

.PHONY: run
run:
	docker-compose exec app go run main.go

.PHONY: test
test:
	docker-compose exec app go test ./...

.PHONY: deep-test
deep-test:
	curl http://127.0.0.1:3001/myself -v

.PHONY: test-elasticsearch
test-elasticsearch:
	curl http://127.0.0.1:9200/_cat/health

.PHONY: search-elasticsearch
search-elasticsearch:
	curl http://127.0.0.1:9200/_search?pretty