package service

import (
	"github.com/Aserose/CaduceusTour/internal/config"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/Aserose/CaduceusTour/internal/repository/mongoDB/data"
	"github.com/Aserose/CaduceusTour/internal/web/client/tg"
	"github.com/Aserose/CaduceusTour/internal/web/sources/gosplan"
	"github.com/Aserose/CaduceusTour/pkg/logger"
)

type Service interface {
	GetDataFromDataSource(params map[string]string) string
	CreateDocument(chatId int64)
	TG
	View
	Overview
}

type TG interface {
	TGMenu(chatId int64, msg string)
	SendMessage(chatId int64, msg string)
}

type View interface {
	Next(chatId int64)
	NextRow(chatId int64)
	ToOverview()
}

type Overview interface {
	General(chatId int64)
	Compare(id int64)
	ToView()
}

type service struct {
	tg           tg.TgApi
	sources      gosplan.GosplanApi
	log          logger.Logger
	strCfg       config.StrConfig
	db           data.OrganizationData
	organization models.Organization
	contracts    []models.ContractInfo
}

func NewService(tg tg.TgApi, sources gosplan.GosplanApi, log logger.Logger, strCfg config.StrConfig) Service {
	return &service{
		tg:      tg,
		sources: sources,
		log:     log,
		strCfg:  strCfg,
	}
}
