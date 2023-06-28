package handler

import (
	"AuthenticationService/pkg/authentication"
	"AuthenticationService/pkg/database"
	"AuthenticationService/pkg/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
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

func (h *Handler) SearchUser(c *gin.Context) {
	var (
		db   database.Database
		resp Response
		data model.TableData
	)
	resp.rw = c.Writer
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}
	db.Connect()
	_, err = db.Access.GetPwd(db.Pool, strconv.FormatInt(data.ChatID, 10))
	if err != nil {
		log.Printf("Don't found password for %d", data.ChatID)
		resp.SetStatusUnauthorized()
	} else {
		resp.SetStatusOk()
	}

}

func (h *Handler) SetUser(c *gin.Context) {
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

	condition := auth.SetPassword(authorization)
	if condition {
		resp.SetStatusOk()
	} else {
		resp.SetStatusUnauthorized()
	}
}
