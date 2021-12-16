package tg

import (
	"github.com/Aserose/CaduceusTour/internal/config"
	"github.com/Aserose/CaduceusTour/pkg/logger"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TgApi interface {
	Init(token, appURL string)
	SendMessage(chatId int64, msgStr string)
	SendFile(chatId int64)
	Menu(chatId int64, msg string)
}

type botApi struct {
	bot    *tgbotapi.BotAPI
	log    logger.Logger
	strCfg config.StrConfig
}

func NewTgApi(log logger.Logger, strCfg config.StrConfig) TgApi {
	return &botApi{
		log:    log,
		strCfg: strCfg,
	}
}
