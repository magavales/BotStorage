package dialog

import (
	"TelegramBot/pkg/model"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type State struct {
	Command string
	Stage   Stage
}

func (s *State) SetState(command string) {
	s.Command = command
	s.Stage = Service
}

func (s *State) HandlerSetStateDialog(message *tgbotapi.Message, dataService *model.DataService) string {
	var response string

	switch s.Stage {
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
		response = "Nice! I save your data!"
		s.clear()
	}

	return response
}

func (s *State) HandlerGetStateDialog(message *tgbotapi.Message, dataService *model.DataService) string {
	var response string

	switch s.Stage {
	case Service:
		dataService.Service = message.Text
		response = fmt.Sprintf("Take your data about service %s", response)
	}

	return response
}

func (s *State) HandlerDelStateDialog(message *tgbotapi.Message, dataService *model.DataService) string {
	var response string

	switch s.Stage {
	case Service:
		dataService.Service = message.Text
		response = fmt.Sprintf("I delete your data about service %s", response)
	}

	return response
}

func (s *State) clear() {
	s.Command = ""
	s.Stage = ""
}
