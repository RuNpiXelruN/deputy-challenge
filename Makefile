.PHONY: up down neo test mocks clearTestCache

mocks: ## generates interface mocks
	go generate ./...

clearTestCache:
	go clean -testcache

up: neo ## Start docker containers
	docker-compose up -d

neo: ## Start Neo4j container
	docker run \
    --name neo4j_deputy \
    -p7474:7474 -p7687:7687 \
	-d \
	--rm \
    -v /${HOME}/neo4j/data:/data \
    -v /${HOME}/neo4j/logs:/logs \
    -v /${HOME}/neo4j/import:/var/lib/neo4j/import \
    -v /${HOME}/neo4j/plugins:/plugins \
    --env NEO4J_AUTH=neo4j/test \
	--env NEO4J_apoc_export_file_enabled=true \
    --env NEO4J_apoc_import_file_enabled=true \
    --env NEO4J_apoc_import_file_use__neo4j__config=true \
    --env NEO4JLABS_PLUGINS=\[\"apoc\"\] \
    neo4j:3.5.13

neo_stop:
	docker stop neo4j_deputy

down: neo_stop ## Stop running docker containers
	docker-compose down

test: clearTestCache ## Run unit tests
	go test -v ./...

build: ## builds the -challenge binary for OSX and places in the root of the workspace
	cd cmd && \
	env GOOS=darwin GOARCH=amd64 go build -v -o deputyJD . && \
	mv deputyJD ../.

buildLinux: ## builds the -challenge binary for Linux and places in the root of the workspace
	cd cmd && \
	env GOOS=linux GOARCH=amd64 go build -v -o deputyJD_linux . && \
	mv deputyJD_linux ../.

buildWin: ## builds the deputy-challenge binary for Windows and places in the root of the workspace
	cd cmd && \
	env GOOS=windows GOARCH=amd64 go build -v -o deputyJD_win . && \
	mv deputyJD_win ../.

help: ## Display available commands	
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'