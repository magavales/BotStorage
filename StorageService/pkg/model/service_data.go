package model

import (
	"encoding/json"
	"io"
	"log"
)

type ServiceData struct {
	ChatID   string `json:"chat_id"`
	Service  string
	Login    string
	Password string
}

func (sd *ServiceData) ParseData(values []interface{}) {
	sd.Service = values[0].(string)
	sd.Login = values[1].(string)
	sd.Password = values[2].(string)
}

func (sd *ServiceData) DecodeJSON(body io.Reader) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&sd)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}
}
