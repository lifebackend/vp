package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/models"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralPostPingHandler Handler for POST /ping
func (h *Handlers) GeneralPostPingHandler(
	params *general.PostPingParams,
	respond *general.PostPingResponses,
) middleware.Responder {
	ctx := context.Background()
	if err := h.authService.Check(ctx, params.Body.ID, params.Body.Password); err != nil {
		return respond.PostPingInternalServerError().FromErr(err)
	}

	ok := models.PostPingResponse{SuccessResponse: models.SuccessResponse{true}}

	return respond.PostPingOK().WithPayload(&ok)
}
