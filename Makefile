deploy-all-local: validate
	docker compose up -d

destroy-all-local:
	docker compose down -v --rmi all --remove-orphans

run-local: deploy-db-local validate serve

deploy-db-local:
	docker compose up -d pack-and-go-mysql

validate: format vet lint test

format:
	go fmt ./...

vet:
	go vet ./...

lint:
	golangci-lint run ./...

test:
	go test -covermode count -coverprofile coverage.out ./...

coverage-local: test
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

serve:
	go run . serve

prepare:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go mod download
	go mod tidy
