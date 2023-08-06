package handlers

import (
	"context"
	"errors"
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
	if params.Body.Body == "" {
		return respond.PostProxyMessageBadRequest().FromErr(errors.New("no content provided for parsing"))
	}

	ctx := context.Background()
	if params.Body.Action != "status" {
		if err := h.authService.Check(ctx, *params.Body.ID, params.HTTPRequest.Header.Get("x-api-key")); err != nil {
			return respond.PostProxyMessageInternalServerError().FromErr(err)
		}
		err := h.messageService.Save(ctx, *params.Body.ID, params.Body.Address, strings.ToLower(params.Body.Action), params.Body.Body)
		if err != nil {
			return respond.PostProxyMessageInternalServerError().FromErr(err)
		}
	}

	ok := models.PostProxyMessagesResponse{models.ProxySuccessResponse{Result: 1}}
	return respond.PostProxyMessageOK().WithPayload(&ok)
}
