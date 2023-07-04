package telegram

import (
	"TelegramBot/pkg/dialog"
	"TelegramBot/pkg/model"
	"TelegramBot/pkg/services/data_service/request"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

var (
	state       dialog.State
	dataService model.DataService
)

type TgBot struct {
	Bot     *tgbotapi.BotAPI
	Uconf   tgbotapi.UpdateConfig
	Updates tgbotapi.UpdatesChannel
}

func (tgb *TgBot) initBot(token string) {
	var err error
	tgb.Bot, err = tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Fatalf("Unable to launch the bot: %s\n", err)
	}
	tgb.Bot.Debug = true
	tgb.Uconf = tgbotapi.NewUpdate(0)
	tgb.Uconf.Timeout = 30
	tgb.Updates, err = tgb.Bot.GetUpdatesChan(tgb.Uconf)
}

func (tgb *TgBot) RunBot(token string) {
	tgb.initBot(token)
	for upd := range tgb.Updates {
		if upd.Message == nil {
			continue
		}

		if upd.Message.IsCommand() {
			tgb.handleCommand(upd.Message)
		} else {
			tgb.handleMessage(upd.Message)
		}
	}
}

func (tgb *TgBot) handleCommand(message *tgbotapi.Message) {
	var (
		req request.Request
	)
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {
	case "start":
		msg.Text = fmt.Sprintf("Hello, %s! I'm BotStorage. I can save your login and password from anything services, if you want.", message.From.UserName)
	case "set":
		if req.SearchUser(msg.ChatID) {
			if model.GetCookie(msg.ChatID) == "" {
				msg.Text = "Write your secret password:"
				state.SetState(message.Command(), dialog.Verify)
			} else {
				msg.Text = "Well! Write name of service:"
				state.SetState(message.Command(), dialog.Service)
			}
		} else {
			msg.Text = "You don't have secret password, so you should set password:"
			state.SetState(message.Command(), dialog.Secret)
		}

		dataService.ChatID = message.Chat.ID
	case "get":
		if req.SearchUser(msg.ChatID) {
			if model.GetCookie(msg.ChatID) == "" {
				msg.Text = "Write your secret password:"
				state.SetState(message.Command(), dialog.Verify)
			} else {
				msg.Text = "Well! Write name of service:"
				state.SetState(message.Command(), dialog.Service)
			}
		} else {
			msg.Text = "You don't have secret password, so you should set password:"
			state.SetState(message.Command(), dialog.Secret)
		}

		dataService.ChatID = message.Chat.ID
	case "del":
		if req.SearchUser(msg.ChatID) {
			if model.GetCookie(msg.ChatID) == "" {
				msg.Text = "Write your secret password:"
				state.SetState(message.Command(), dialog.Verify)
			} else {
				msg.Text = "Well! Write name of service:"
				state.SetState(message.Command(), dialog.Service)
			}
		} else {
			msg.Text = "You don't have secret password, so you should set password:"
			state.SetState(message.Command(), dialog.Secret)
		}

		dataService.ChatID = message.Chat.ID
	case "secret":
		msg.Text = fmt.Sprintf("I love thee very much, baby!❤️\nWhat about you?")
	}

	_, _ = tgb.Bot.Send(msg)
}

func (tgb *TgBot) handleMessage(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch state.Command {
	case "set":
		msg.Text = state.HandlerSetStateDialog(message, &dataService)
	case "get":
		msg.Text = state.HandlerGetStateDialog(message, &dataService)
	case "del":
		msg.Text = state.HandlerDelStateDialog(message, &dataService)
	}

	_, _ = tgb.Bot.Send(msg)
}
