.PHONY: run swagger test-unit-one test-unit-all test-endpoint-one test-endpoint-all mock-service-one mock-service-all test-code-coverage
# make run
run:
	go run cmd/api/main.go
swagger:
	swag init -g cmd/api/main.go
# make test-unit-one t=<Name of function to test> path=<package-path>(e.g., ./internal/service)
GO-TEST-ARGS:=go test -v -cover
test-unit-one:
	$(if $(t),,$(error missing t))
	$(if $(path),,$(error missing path))
	$(GO-TEST-ARGS) -run $(t) $(path)
# make test-unit-all path=<package-path>(e.g., ./internal/service/...)
test-unit-all:
	$(if $(path),,$(error missing path))
	$(GO-TEST-ARGS) $(path)
	
# make test-endpoint-one t=<Name of function to test>
PATH-TEST-ENDPOINT:=./internal/test/endpoint
test-endpoint-one:
	$(if $(t),,$(error missing t))
	$(GO-TEST-ARGS) -run $(t) $(PATH-TEST-ENDPOINT)
# make test-endpoint-all
test-endpoint-all:
	$(GO-TEST-ARGS) $(PATH-TEST-ENDPOINT)/...

# make test-code-coverage
test-code-coverage:
	go test ./... -coverprofile=coverage.tmp -covermode=atomic -coverpkg=./... -p 1
	go run filter_coverage.go
	go tool cover -html=coverage.out -o coverage.html

# make mock-one folder=<package-folder-for-mock>
mock-one:
	$(if $(folder),,$(error missing folder))
	cd $(folder) && go generate ./...
# make mock-service-all
mock-service-all:
	go generate ./...