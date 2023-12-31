package handler

import (
	"StorageService/pkg/database"
	"StorageService/pkg/model"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) GetData(c *gin.Context) {
	var (
		resp        Response
		serviceData model.ServiceData
		db          database.Database
	)
	log.Println("Request has been gotten!")
	resp.rw = c.Writer
	serviceData.DecodeJSON(c.Request.Body)

	db.Connect()
	data, err := db.Access.GetData(db.Pool, serviceData.ChatID, serviceData.Service)
	if err != nil {
		log.Printf("Can't get data from table: %s\n", err)
		resp.SetStatusBadRequest()
		log.Println("Response has been sent!")
		return
	} else {
		jdata, _ := json.Marshal(data)
		resp.SetData(jdata)
		resp.SetStatusOk()
		log.Println("Response has been sent!")
	}
}

func (h *Handler) AddData(c *gin.Context) {
	var (
		resp        Response
		serviceData model.ServiceData
		db          database.Database
	)
	log.Println("Request has been gotten!")
	resp.rw = c.Writer
	serviceData.DecodeJSON(c.Request.Body)

	db.Connect()
	err := db.Access.AddData(db.Pool, serviceData)
	if err != nil {
		log.Printf("Unpossible add data: %s\n", err)
		resp.SetStatusBadRequest()
		log.Println("Response has been sent!")
		return
	} else {
		resp.SetStatusOk()
		log.Println("Response has been sent!")
	}
}

func (h *Handler) DelData(c *gin.Context) {
	var (
		resp        Response
		serviceData model.ServiceData
		db          database.Database
	)
	log.Println("Request has been gotten!")
	resp.rw = c.Writer
	serviceData.DecodeJSON(c.Request.Body)

	db.Connect()
	err := db.Access.DelData(db.Pool, serviceData)
	if err != nil {
		log.Printf("Can't del data from table: %s\n", err)
		resp.SetStatusBadRequest()
		log.Println("Response has been sent!")
		return
	} else {
		resp.SetStatusOk()
		log.Println("Response has been sent!")
	}
}
