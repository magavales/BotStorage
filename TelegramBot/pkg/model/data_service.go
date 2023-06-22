package model

import (
	"encoding/json"
	"io"
	"log"
)

type DataService struct {
	ChatID   int64  `json:"chat_id"`
	Service  string `json:"service"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (ds *DataService) Clear() {
	ds.ChatID = 0
	ds.Service = ""
	ds.Login = ""
	ds.Password = ""
}

func (ds *DataService) ToJSON() []byte {
	body, err := json.Marshal(ds)
	if err != nil {
		log.Printf("Convertation has been stopped!: %s", err)
	}

	return body
}

func (ds *DataService) DecodeJSON(body io.Reader) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&ds)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}
}
