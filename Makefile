DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASS=admin
PORT=8080
DB_CRED=DB_HOST=$(DB_HOST) DB_PORT=$(DB_PORT) DB_USER=$(DB_USER) DB_PASS=$(DB_PASS)

run-cli-memory:
	cd cmd/translator-memory/cli; go build -v -o translator-memory-cli -mod vendor -race -tags='cli memory'
	./cmd/translator-memory/cli/translator-memory-cli -race

run-web-memory:
	cd cmd/translator-memory/web; go build -v -o translator-memory-web -mod vendor -race -tags='memory web'
	./cmd/translator-memory/web/translator-memory-web -p=$(PORT)

run-cli-db:
	cd cmd/translator-db/cli; go build -v -o translator-db-cli -mod vendor -race -tags='db cli'
	$(DB_CRED) ./cmd/translator-db/cli/translator-db-cli -race

run-web-db:
	cd cmd/translator-db/web; go build -v -o translator-db-web -mod vendor -race -tags='db web'
	$(DB_CRED) ./cmd/translator-db/web/translator-db-web -p=$(PORT)

test-us-cli-mem:
	go test ./internal/test/userstory -run TestMemCLI -count=1 -v -race -tags='memory db cli web'

test-us-cli-db:
	$(DB_CRED) go test ./internal/test/userstory -run TestDBCLI -count=1 -v -race -tags='memory db cli web'

test-us-web-mem:
	go test ./internal/test/userstory -run TestMemWeb -count=1 -v -race -tags='memory db cli web'

test-us-web-db:
	$(DB_CRED) go test ./internal/test/userstory -run TestDBWeb -count=1 -v -race -tags='memory db cli web'

test-all:
	$(DB_CRED) go test {./internal/dao,./internal/test/...} -count=1 -tags='memory db cli web'