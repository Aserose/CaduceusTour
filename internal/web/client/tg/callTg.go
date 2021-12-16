package tg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	AllRegion = "РФ"
)

func (b *botApi) SendMessage(chatId int64, msgStr string) {
	msg := tgbotapi.NewMessage(chatId, msgStr)
	_, err := b.bot.Send(msg)
	if err != nil {
		b.log.Error("tg: send message error: %v",err.Error())
	}
}

func (b *botApi) SendFile(chatId int64) {
	msg := tgbotapi.NewDocumentUpload(chatId, "result.txt")

	_, err := b.bot.Send(msg)
	if err != nil {
		b.log.Error("tg: send document error: %v",err.Error())
	}
}

func (b *botApi) Menu(chatId int64, msg string) {

	menu := tgbotapi.NewMessage(chatId, "_")

	menuButton := map[string][]string{
		b.strCfg.Menu.Button.Menu:         {b.strCfg.Menu.Button.Organization, b.strCfg.Menu.Button.Region},
		b.strCfg.Menu.Button.Region:       {AllRegion, b.strCfg.Menu.Button.Code, b.strCfg.Menu.Button.Back},
		b.strCfg.Menu.Button.Organization: {b.strCfg.Menu.Button.Name, b.strCfg.Menu.Button.Address, b.strCfg.Menu.Button.Back},
		b.strCfg.Menu.Button.Intermediate: {b.strCfg.Menu.Button.Overview, b.strCfg.Menu.Button.List, b.strCfg.Menu.Button.Back},
		b.strCfg.Menu.Button.Overview:     {b.strCfg.Menu.Button.Compare, b.strCfg.Menu.Button.List, b.strCfg.Menu.Button.Back},
		b.strCfg.Menu.Button.View: {b.strCfg.Menu.Button.Next, b.strCfg.Menu.Button.Load,
			b.strCfg.Menu.Button.NextRow, b.strCfg.Menu.Button.Overview, b.strCfg.Menu.Button.Back},
	}

	button := tgbotapi.NewKeyboardButtonRow()

	for _, menuButton := range menuButton[msg] {
		button = append(button, tgbotapi.NewKeyboardButton(menuButton))
	}

	menu.ReplyMarkup = tgbotapi.NewReplyKeyboard(button)
	_, err := b.bot.Send(menu)
	if err != nil {
		b.log.Errorf("tg: menu creation error, %v", err.Error())
	}
}
