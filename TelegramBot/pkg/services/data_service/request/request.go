package request

import (
	"TelegramBot/pkg/model"
	"TelegramBot/pkg/services/data_service"
	"TelegramBot/pkg/services/data_service/response"
	"bytes"
	"encoding/base64"
	"log"
	"net/http"
	"strconv"
)

type Request struct {
	Request *http.Request
}

func (r *Request) SetData(dataService *model.DataService) bool {
	var err error
	body := dataService.ToJSON()
	r.Request, err = http.NewRequest("POST", "http://localhost:8080/api/v1/add", bytes.NewReader(body))
	if err != nil {
		log.Println("Invalid request!")
	}
	r.Request.Header.Set("Content-Type", "application/json")
	resp := r.doRequest()
	if resp.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}

func (r *Request) GetData(dataService *model.DataService) (model.DataService, bool) {
	var (
		resp response.Response
		data model.DataService
		err  error
	)
	body := dataService.ToJSON()
	r.Request, err = http.NewRequest("GET", "http://localhost:8080/api/v1/get", bytes.NewReader(body))
	if err != nil {
		log.Println("Invalid request!")
	}
	r.Request.Header.Set("Content-Type", "application/json")
	resp.Response = r.doRequest()
	if resp.Response.StatusCode == http.StatusOK {
		log.Println(resp.Response.Body)
		data.DecodeJSON(resp.Response.Body)
		return data, true
	} else {
		return data, false
	}
}

func (r *Request) DelData(dataService *model.DataService) bool {
	var (
		err  error
		resp response.Response
	)
	body := dataService.ToJSON()
	r.Request, err = http.NewRequest("GET", "http://localhost:8080/api/v1/del", bytes.NewReader(body))
	if err != nil {
		log.Println("Invalid request!")
	}
	r.Request.Header.Set("Content-Type", "application/json")
	resp.Response = r.doRequest()
	if resp.Response.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}

func (r *Request) SearchUser(chatID int64) bool {
	var (
		err         error
		resp        response.Response
		dataService model.DataService
	)
	dataService.ChatID = chatID
	body := dataService.ToJSON()
	r.Request, err = http.NewRequest("GET", "http://localhost:8080/api/v1/search", bytes.NewReader(body))
	if err != nil {
		log.Println("Invalid request!")
	}
	r.Request.Header.Set("Content-Type", "application/json")
	resp.Response = r.doRequest()
	if resp.Response.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}

func (r *Request) VerifyUser(chatID int64, text string) bool {
	var (
		err     error
		encoded []byte
		resp    response.Response
	)
	body := strconv.FormatInt(chatID, 10) + ":" + text
	encoded = []byte(base64.StdEncoding.EncodeToString([]byte(body)))
	r.Request, err = http.NewRequest("POST", "http://localhost:8080/api/v1/auth", bytes.NewReader([]byte("")))
	if err != nil {
		log.Println("Invalid request!")
	}
	r.Request.Header.Set("Authorization", "Basic "+string(encoded))
	resp.Response = r.doRequest()
	if resp.Response.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}

func (r *Request) SetUser(chatID int64, text string) bool {
	var (
		err     error
		encoded []byte
		resp    response.Response
	)
	body := strconv.FormatInt(chatID, 10) + ":" + text
	encoded = []byte(base64.StdEncoding.EncodeToString([]byte(body)))
	r.Request, err = http.NewRequest("POST", "http://localhost:8080/api/v1/set", bytes.NewReader([]byte("")))
	if err != nil {
		log.Println("Invalid request!")
	}
	r.Request.Header.Set("Authorization", "Basic "+string(encoded))
	resp.Response = r.doRequest()
	if resp.Response.StatusCode == http.StatusOK {
		return true
	} else {
		return false
	}
}

func (r *Request) doRequest() *http.Response {
	client := new(data_service.Client).InitClient()
	resp, err := client.Do(r.Request)
	if err != nil {
		log.Printf("Client hasn't been do smth!: %s\n", err)
	}

	return resp
}
