go install github.com/vektra/mockery/v2@latest

## 🔧 Generate mocks (mockery)
 //go:generate mockery --name GenPassword --filename gen_password_mock.go --output ./mocks
 //go:generate mockery --name HealthCheck --filename check_health_mock.go --output ./mocks

```bash
go generate ./internal/service