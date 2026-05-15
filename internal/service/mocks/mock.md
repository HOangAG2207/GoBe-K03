go install github.com/vektra/mockery/v2@latest

## 🔧 Generate mocks (mockery) - import this code to in front of service interface
 //go:generate mockery --name GenPassword --filename gen_password_mock.go --output ./mocks
 //go:generate mockery --name HealthCheck --filename check_health_mock.go --output ./mocks

## make file run
make mock-service-one name=GenPassword file=gen_password_mock.go
make mock-service-one name=HealthCheck file=check_health_mock.go

```bash
go generate ./internal/service