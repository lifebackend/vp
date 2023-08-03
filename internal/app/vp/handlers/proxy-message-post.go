package handlers

import (
	"context"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/models"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
)

// GeneralPostProxyMessageHandler Handler for POST /proxy/message
func (h *Handlers) GeneralPostProxyMessageHandler(
	params *general.PostProxyMessageParams,
	respond *general.PostProxyMessageResponses,
) middleware.Responder {

	ctx := context.Background()
	if err := h.authService.Check(ctx, *params.Body.ID, params.HTTPRequest.Header.Get("x-api-key")); err != nil {
		return respond.PostProxyMessageInternalServerError().FromErr(err)
	}
	err := h.messageService.Save(ctx, *params.Body.ID, params.Body.Address, strings.ToLower(params.Body.Action), params.Body.Body)
	if err != nil {
		return respond.PostProxyMessageInternalServerError().FromErr(err)
	}

	ok := models.PostProxyMessagesResponse{models.ProxySuccessResponse{Result: 1}}
	return respond.PostProxyMessageOK().WithPayload(&ok)
}
