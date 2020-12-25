.PHONY: test
test:
	clear
	go test -count=1 -timeout 30s -p 1 -v -cover ./...

.PHONY: cover
cover:
	clear
	go test -count=1 -timeout 60s -p 1 -coverprofile=./cover/all-profile.out -coverpkg=./... ./...
	go tool cover -html=./cover/all-profile.out -o ./cover/all-coverage.html

lint:
	golangci-lint run ./...
