package handlers

import "github.com/lifebackend/vp/internal/app/vp/message"

func NewHandlers(
	imageTag string,
	messageService *message.Service,
) *Handlers {
	return &Handlers{
		imageTag:       imageTag,
		messageService: messageService,
	}
}

type Handlers struct {
	imageTag       string
	messageService *message.Service
}
