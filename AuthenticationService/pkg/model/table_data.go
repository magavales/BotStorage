package model

type TableData struct {
	ChatID int64 `json:"chat_id"`
	Pwd    string
	Salt   string
}

func (td *TableData) ParseData(values []interface{}) {
	td.Pwd = values[0].(string)
	td.Salt = values[1].(string)
}
