run-cli-memory:
	cd cmd/translator-memory/cli; go build -v -o translator-memory-cli -mod vendor -race
	./cmd/translator-memory/cli/translator-memory-cli -race

run-web-memory:
	cd cmd/translator-memory/web; go build -v -o translator-memory-web -mod vendor -race
	./cmd/translator-memory/web/translator-memory-web -p=8080