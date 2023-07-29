package handlers

import (
	"context"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/models"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralPostSmsHandler Handler for POST /sms
func (h *Handlers) GeneralPostSmsHandler(
	params *general.PostSmsParams,
	respond *general.PostSmsResponses,
) middleware.Responder {
	ctx := context.Background()
	if err := h.authService.Check(ctx, params.Body.ID, params.Body.Password); err != nil {
		return respond.PostSmsInternalServerError().FromErr(err)
	}
	err := h.messageService.Save(ctx, params.Body.ID, params.Body.From, "sms", params.Body.Message)
	if err != nil {
		return respond.PostSmsInternalServerError().FromErr(err)
	}
	ok := models.PostMessageResponse{SuccessResponse: models.SuccessResponse{true}}

	return respond.PostSmsOK().WithPayload(&ok)
}
