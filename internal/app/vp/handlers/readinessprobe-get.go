package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/health"
)

// HealthGetReadinessProbeHandler Handler for GET /_readinessProbe
func (h *Handlers) HealthGetReadinessProbeHandler(
	params *health.GetReadinessProbeParams,
	respond *health.GetReadinessProbeResponses,
) middleware.Responder {
	return middleware.NotImplemented("operation health.GetReadinessProbe has not yet been implemented")
}
