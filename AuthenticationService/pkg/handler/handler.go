package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouter() *gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		version := api.Group("/v1")
		{
			version.POST("/auth", h.Authentication)
		}
	}

	return router
}
