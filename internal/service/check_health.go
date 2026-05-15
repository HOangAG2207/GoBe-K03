// service/check_health.go
package service

import "github.com/google/uuid"

// HealthCheck defines the contract for health check business logic
type HealthCheck interface {
	CheckHealth() HealthCheckResponse
}

// healthCheckService implements HealthCheck interface
// It holds service metadata such as service name and instance ID
type healthCheckService struct {
	serviceName string
	instanceID  string
}

// HealthCheckResponse represents the response returned by the health check API
type HealthCheckResponse struct {
	Message     string `json:"message"`      // status message (e.g., "OK")
	ServiceName string `json:"service_name"` // name of the service
	InstanceID  string `json:"instance_id"`  // unique instance identifier
}

// NewHealthCheck creates a new HealthCheck service instance
// It supports dependency injection of service name and instance ID
func NewHealthCheck(serviceName, instanceID string) HealthCheck {
	return &healthCheckService{
		serviceName: serviceName,
		instanceID:  instanceID,
	}
}

// CheckHealth returns the health status of the service
// It ensures that an instance ID always exists (generates UUID if missing)
func (h *healthCheckService) CheckHealth() HealthCheckResponse {
	instanceID := h.instanceID

	// Generate a new UUID if instanceID is not provided
	if instanceID == "" {
		instanceID = uuid.New().String()
	}

	// Return standardized health check response
	return HealthCheckResponse{
		Message:     "OK",
		ServiceName: h.serviceName,
		InstanceID:  instanceID,
	}
}
