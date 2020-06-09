.PHONY: test
test:
	clear
	go test -count=1 -timeout 10s -p 1 -v ./...

test_infrastructure_cover:
	clear
	go test -count=1 -timeout 10s -p 1 -coverprofile=./cover/infrastructure-profile.out -covermode=atomic ./infrastructure
	go tool cover -html=./cover/infrastructure-profile.out -o ./cover/infrastructure-coverage.html
