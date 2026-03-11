# go

lint:
	golangci-lint run ./... -c=golangci-lint.yml

test:
	go test -parallel=1 -count=1 ./... -coverprofile=cover.out
	go tool cover -func=cover.out | tail -n1

run:
	go run cmd/main.go
