package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralGetAppCodesHandler Handler for GET /app-codes
func (h *Handlers) GeneralGetAppCodesHandler(
	params *general.GetAppCodesParams,
	respond *general.GetAppCodesResponses,
) middleware.Responder {
	return middleware.NotImplemented("operation general.GetAppCodes has not yet been implemented")
}
