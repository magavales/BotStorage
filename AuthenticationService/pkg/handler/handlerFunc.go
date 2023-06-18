package handler

import (
	"AuthenticationService/pkg/authentication"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) Authentication(c *gin.Context) {
	var (
		auth authentication.Authentication
		resp Response
	)
	resp.rw = c.Writer
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		resp.SetStatusBadRequest()
		log.Println("Bad request!")
	}

	condition := auth.VerificationPassword(authorization)
	if condition {
		resp.SetStatusOk()
	} else {
		resp.SetStatusUnauthorized()
	}

}
