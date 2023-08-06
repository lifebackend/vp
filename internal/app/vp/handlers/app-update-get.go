package handlers

import (
	"context"
	"errors"
	"os"

	"github.com/go-openapi/runtime/middleware"
	"github.com/lifebackend/vp/internal/app/vp/server/restapi/operations/general"
	"go.mongodb.org/mongo-driver/bson"
)

// GeneralGetAppUpdateHandler Handler for GET /app/update
func (h *Handlers) GeneralGetAppUpdateHandler(
	params *general.GetAppUpdateVersionParams,
	respond *general.GetAppUpdateVersionResponses,
) middleware.Responder {
	version := struct {
		Version string `bson:"version"`
		Path    string `bson:"path"`
	}{}

	filter := bson.D{{"version", bson.D{{"$ne", params.Version}}}}
	err := h.mongoClient.Database("database").
		Collection("version").
		FindOne(context.Background(), filter).Decode(&version)
	if err != nil {
		return respond.GetAppUpdateVersionInternalServerError().FromErr(err)
	}

	if version.Path == "" {
		return respond.GetAppUpdateVersionInternalServerError().FromErr(errors.New("apk not found"))
	}

	f, err := os.Open(version.Path)
	if err != nil {
		return respond.GetAppUpdateVersionInternalServerError().FromErr(err)
	}

	return respond.GetAppUpdateVersionOK().WithPayload(f)
}
