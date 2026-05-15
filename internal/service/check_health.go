// service/check_health.go
package service

import "github.com/google/uuid"

type HealthCheck interface {
	CheckHealth() HealthCheckResponse
}

type healthCheckService struct {
	serviceName string
	instanceID  string
}

type HealthCheckResponse struct {
	Message     string `json:"message"`
	ServiceName string `json:"service_name"`
	InstanceID  string `json:"instance_id"`
}

func NewHealthCheck(serviceName, instanceID string) HealthCheck {
	return &healthCheckService{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}
func (h *healthCheckService) CheckHealth() HealthCheckResponse {
	instanceID := h.instanceID

	if instanceID == "" {
		instanceID = uuid.New().String()
	}
	return HealthCheckResponse{
		Message:     "OK",
		ServiceName: h.serviceName,
		InstanceID:  instanceID,
	}
}
