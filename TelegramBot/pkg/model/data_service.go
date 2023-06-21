package model

type DataService struct {
	ChatID   int64
	Service  string
	Login    string
	Password string
}

func (ds *DataService) Clear() {
	ds.ChatID = 0
	ds.Service = ""
	ds.Login = ""
	ds.Password = ""
}
