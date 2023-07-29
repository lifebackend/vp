package handlers

import (
	"github.com/lifebackend/vp/internal/app/vp/auth"
	"github.com/lifebackend/vp/internal/app/vp/message"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewHandlers(
	imageTag string,
	messageService *message.Service,
	a *auth.Service,
	client *mongo.Client,
) *Handlers {
	return &Handlers{
		imageTag:       imageTag,
		messageService: messageService,
		authService:    a,
		mongoClient:    client,
	}
}

type Handlers struct {
	imageTag       string
	messageService *message.Service
	authService    *auth.Service
	mongoClient    *mongo.Client
}
