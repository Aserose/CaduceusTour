package tg

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func (b *botApi) Init(token, appURL string) {
	b.log.Info("tg: initialization")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		b.log.Errorf("tg: %v", err.Error())
	}
	b.bot = bot

	b.bot.Debug = true

	b.log.Infof("Authorized on account %s", bot.Self.UserName)

	b.setWebhook(appURL)
}

func (b *botApi) setWebhook(appURL string) {
	b.log.Info("tg: set webhook")
	_, err := b.bot.SetWebhook(tgbotapi.NewWebhook(appURL))
	if err != nil {
		b.log.Errorf("tg: create webhook error: %v", err.Error())
	}
	info, err := b.bot.GetWebhookInfo()
	if err != nil {
		b.log.Errorf("tg: get webhook error: %v", err.Error())
	}
	if info.LastErrorDate != 0 {
		b.log.Printf("Telegram callback failed: %s", info.LastErrorMessage)
	}
	b.log.Info("tg: webhook is set")
}
