package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/models"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralPostPushHandler Handler for POST /push
func (h *Handlers) GeneralPostPushHandler(
	params *general.PostPushParams,
	respond *general.PostPushResponses,
) middleware.Responder {

	ctx := context.Background()
	if err := h.authService.Check(ctx, *params.Body.ID, params.Body.Password); err != nil {
		return respond.PostPushInternalServerError().FromErr(err)
	}
	err := h.messageService.Save(ctx, *params.Body.ID, params.Body.From, "push", params.Body.Message)
	if err != nil {
		return respond.PostPushInternalServerError().FromErr(err)
	}

	ok := models.PostMessageResponse{models.SuccessResponse{Success: true}}
	return respond.PostPushOK().WithPayload(&ok)
}
