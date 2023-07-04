package dialog

import (
	"TelegramBot/pkg/model"
	"TelegramBot/pkg/services/data_service/request"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type State struct {
	Command string
	Stage   Stage
}

func (s *State) SetState(command string, stage Stage) {
	s.Command = command
	s.Stage = stage
}

func (s *State) HandlerSetStateDialog(message *tgbotapi.Message, dataService *model.DataService) string {
	var (
		response string
		req      request.Request
	)

	switch s.Stage {
	case Secret:
		if req.SetUser(message.Chat.ID, message.Text) {
			s.Stage = Service
			response = fmt.Sprintf("Well! Write name of service:")
		} else {
			response = fmt.Sprintf("Invalid password!")
			s.clear()
			dataService.Clear()
		}
	case Verify:
		if req.VerifyUser(message.Chat.ID, message.Text) {
			s.Stage = Service
			response = fmt.Sprintf("Well! Write name of service:")
		} else {
			response = fmt.Sprintf("Invalid password!")
			s.clear()
			dataService.Clear()
		}
	case Service:
		dataService.Service = message.Text
		s.Stage = Login
		response = "Ok, go next! Write your login:"
	case Login:
		dataService.Login = message.Text
		s.Stage = Password
		response = "Well done! Write your password:"
	case Password:
		dataService.Password = message.Text
		if req.SetData(dataService) {
			response = "Nice! I save your data!"
		} else {
			response = "I can't save your data!"
		}
		s.clear()
		dataService.Clear()
	}

	return response
}

func (s *State) HandlerGetStateDialog(message *tgbotapi.Message, dataService *model.DataService) string {
	var (
		response string
		req      request.Request
	)

	switch s.Stage {
	case Secret:
		if req.SetUser(message.Chat.ID, message.Text) {
			s.Stage = Service
			response = fmt.Sprintf("Well! Write name of service:")
		} else {
			response = fmt.Sprintf("Invalid password!")
			s.clear()
			dataService.Clear()
		}
	case Verify:
		if req.VerifyUser(message.Chat.ID, message.Text) {
			s.Stage = Service
			response = fmt.Sprintf("Well! Write name of service:")
		} else {
			response = fmt.Sprintf("Invalid password!")
			s.clear()
			dataService.Clear()
		}
	case Service:
		dataService.Service = message.Text
		data, condition := req.GetData(dataService)
		if condition {
			response = fmt.Sprintf("Take your data about service %s.\nLogin: %s\nPassword: %s", data.Service, data.Login, data.Password)
		} else {
			response = fmt.Sprintf("I can't take your data about service %s.", data.Service)
		}

	}

	return response
}

func (s *State) HandlerDelStateDialog(message *tgbotapi.Message, dataService *model.DataService) string {
	var (
		response string
		req      request.Request
	)

	switch s.Stage {
	case Secret:
		if req.SetUser(message.Chat.ID, message.Text) {
			s.Stage = Service
			response = fmt.Sprintf("Well! Write name of service:")
		} else {
			response = fmt.Sprintf("Invalid password!")
			s.clear()
			dataService.Clear()
		}
	case Verify:
		if req.VerifyUser(message.Chat.ID, message.Text) {
			s.Stage = Service
			response = fmt.Sprintf("Well! Write name of service:")
		} else {
			response = fmt.Sprintf("Invalid password!")
			s.clear()
			dataService.Clear()
		}
	case Service:
		dataService.Service = message.Text
		if req.DelData(dataService) {
			response = fmt.Sprintf("I delete your data about service %s", message.Text)
		} else {
			response = fmt.Sprintf("I can't delete your data about service %s", message.Text)
		}

	}

	return response
}

func (s *State) clear() {
	s.Command = ""
	s.Stage = ""
}
