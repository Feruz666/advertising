package handler

import (
	"context"

	"github.com/Feruz666/advertising/pkg/service"
	"github.com/labstack/echo/v4"
)

type PostHandler struct {
	ctx context.Context
	services *service.Manager
}


func NewPost(ctx context.Context, services *service.Manager) *PostHandler {
	return &PostHandler{
		ctx: ctx,
		services: services,
	}
}

func (h *PostHandler) Get(ctx echo.Context) error {
	return nil
}
