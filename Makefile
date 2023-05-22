.PHONY: lint test vendor clean

default: lint test

lint:
	golangci-lint run

test:
	go test -v -cover ./...

yaegi_test:
	yaegi test -v .

vendor:
	go mod vendor

entr:
	# https://github.com/eradman/entr
	find | entr -r -s "docker compose up --remove-orphans"

clean:
	rm -rf ./vendor