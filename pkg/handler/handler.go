package handler

import (
	"github.com/Feruz666/advertising/pkg/service"
)

type Handler struct {
	services *service.Manager
}


func NewHandler(services *service.Manager) *Handler {
	return &Handler{
		services: services,
	}
}

// func (h *Handler) InitRoutes() *echo.Echo {
// 	router := echo.New()

// 	api := router.Group("/api")
// 	adverts := api.Group("/adverts")
// 	// adverts.GET("", h.findAll)

// 	return router
// }