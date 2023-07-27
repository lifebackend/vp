package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralPostSmsHandler Handler for POST /sms
func (h *Handlers) GeneralPostSmsHandler(
	params *general.PostSmsParams,
	respond *general.PostSmsResponses,
) middleware.Responder {
	return middleware.NotImplemented("operation general.PostSms has not yet been implemented")
}
