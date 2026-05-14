.PHONY: run test-service-one GoBe-K03

# make run
run:
	go run cmd/api/main.go

# make test-service-one t=<Name of function to test>
test-service-one:
	go test -v -cover -run $(t) ./internal/service
# make test-service-all
test-service-all:
	go test -v -cover ./internal/service/...
# make mock-one name=<service-interface-name> file=<go-filename-for-mock>
mock-one:
	cd internal/service && \
	mockery --name $(name) --filename $(file) --output ./mocks
# make mock-service-all
mock-service-all:
	go generate ./internal/service