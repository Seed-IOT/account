.PHONY: run
run:
	go run cmd/main.go

.PHONY: test
test:
	cd cmd;go test

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down