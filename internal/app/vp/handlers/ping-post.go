package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralPostPingHandler Handler for POST /ping
func (h *Handlers) GeneralPostPingHandler(
	params *general.PostPingParams,
	respond *general.PostPingResponses,
) middleware.Responder {
	return middleware.NotImplemented("operation general.PostPing has not yet been implemented")
}
