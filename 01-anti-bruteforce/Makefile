help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-16s\033[0m %s\n", $$1, $$2}'
.PHONY: help

run: ## docker up
	docker-compose up -d
.PHONY: run

build: ## build service anti bruteforce
	go build -o server .
.PHONY: build

docker.rebuild: ## docker rebuild
	docker-compose up -d --build
.PHONY: docker.rebuild

down: ## docker down
	docker-compose down --remove-orphans

test: ## run test
	go test tests/anti-bruteforce_test.go
