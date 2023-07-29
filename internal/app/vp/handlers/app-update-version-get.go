package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralGetAppUpdateVersionHandler Handler for GET /app/update/{version}
func (h *Handlers) GeneralGetAppUpdateVersionHandler(
	params *general.GetAppUpdateVersionParams,
	respond *general.GetAppUpdateVersionResponses,
) middleware.Responder {
	return middleware.NotImplemented("operation general.GetAppUpdateVersion has not yet been implemented")
}
