package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralGetAppUpdateHandler Handler for GET /app/update
func (h *Handlers) GeneralGetAppUpdateHandler(
	params *general.GetAppUpdateParams,
	respond *general.GetAppUpdateResponses,
) middleware.Responder {
	return middleware.NotImplemented("operation general.GetAppUpdate has not yet been implemented")
}
