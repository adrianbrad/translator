run-cli-memory:
	cd cmd/translator-memory/cli; go build -v -o translator-memory-cli -mod vendor -race -tags='cli memory'
	./cmd/translator-memory/cli/translator-memory-cli -race

run-web-memory:
	cd cmd/translator-memory/web; go build -v -o translator-memory-web -mod vendor -race -tags='memory web'
	./cmd/translator-memory/web/translator-memory-web -p=8080

run-cli-db:
	cd cmd/translator-db/cli; go build -v -o translator-db-cli -mod vendor -race -tags='db cli'
	DB_HOST=localhost DB_PORT=5432 DB_USER=admin DB_PASS=admin ./cmd/translator-db/cli/translator-db-cli -race

run-web-db:
	cd cmd/translator-db/web; go build -v -o translator-db-web -mod vendor -race -tags='db web'
	DB_HOST=localhost DB_PORT=5432 DB_USER=admin DB_PASS=admin ./cmd/translator-db/web/translator-db-web -p=8080