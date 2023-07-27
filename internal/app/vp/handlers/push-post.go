package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralPostPushHandler Handler for POST /push
func (h *Handlers) GeneralPostPushHandler(
	params *general.PostPushParams,
	respond *general.PostPushResponses,
) middleware.Responder {
	return middleware.NotImplemented("operation general.PostPush has not yet been implemented")
}
