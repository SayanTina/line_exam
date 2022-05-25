package models

import "time"

type HealthCheckResp struct {
	ServiceAmount     int           `json:"service_amount"`
	UpServiceAmount   int           `json:"up_service_amount"`
	DownServiceAmount int           `json:"down_service_amount"`
	ExecutionTime     time.Duration `json:"execution_time"`
}
