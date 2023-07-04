package model

import (
	"encoding/json"
	"io"
	"log"
)

type TableData struct {
	ChatID int64 `json:"chat_id"`
	Pwd    string
	Salt   string
}

func (td *TableData) ParseData(values []interface{}) {
	td.Pwd = values[0].(string)
	td.Salt = values[1].(string)
}

func (td *TableData) DecodeData(body io.ReadCloser) {
	decoder := json.NewDecoder(body)
	err := decoder.Decode(&td.ChatID)
	if err != nil {
		log.Println("JSON hasn't been decoded!")
	}
}
