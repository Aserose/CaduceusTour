package middleware

import (
	"github.com/Aserose/CaduceusTour/internal/config"
	"github.com/Aserose/CaduceusTour/internal/service"
	"github.com/Aserose/CaduceusTour/pkg/logger"
)

type Middleware interface {
	Begin(msg string, id int64, menu Menu)
	Request(msg string, id int64, menu Menu)
	CreateMenu() Menu
}

type middleware struct {
	service service.Service
	log     logger.Logger
	strCfg  config.StrConfig
	cfg     *config.HDLConfig
}

func NewHandler(service service.Service, log logger.Logger, strCfg config.StrConfig, cfg *config.HDLConfig) Middleware {
	return &middleware{
		service: service,
		log:     log,
		strCfg:  strCfg,
		cfg:     cfg,
	}
}
