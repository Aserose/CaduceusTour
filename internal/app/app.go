package app

import (
	"context"
	"github.com/Aserose/CaduceusTour/internal/config"
	"github.com/Aserose/CaduceusTour/internal/middleware"
	"github.com/Aserose/CaduceusTour/internal/repository/mongoDB"
	"github.com/Aserose/CaduceusTour/internal/repository/mongoDB/data"
	"github.com/Aserose/CaduceusTour/internal/server"
	"github.com/Aserose/CaduceusTour/internal/service"
	"github.com/Aserose/CaduceusTour/internal/web/client/tg"
	"github.com/Aserose/CaduceusTour/internal/web/sources/gosplan"
	"github.com/Aserose/CaduceusTour/pkg/logger"
)

func Start() {
	log := logger.NewLogger()

	cfgTG, cfgDB, cfgGP, cfgHDL, err := config.Init()
	if err != nil {
		log.Error(err.Error())
	}

	db, ctxDB, err := mongoDB.MongoInit(context.Background(), cfgDB.MongoDB.Host, cfgDB.MongoDB.Port,
		cfgDB.MongoDB.Username, cfgDB.MongoDB.Password, cfgDB.MongoDB.DBName, log)
	if err != nil {
		log.Error(err.Error())
	}

	strCfg, transcription, serverCfg := config.InitStrCfg("configs/configs.yml")

	gpData := data.NewMongoData(db, db.Collection(cfgDB.MongoDB.GPData), db.Collection(cfgDB.MongoDB.Access), ctxDB, log)

	tgApi := tg.NewTgApi(log, strCfg)
	tgApi.Init(cfgTG.TokenTG, cfgTG.AppURL)
	gosplanApi := gosplan.NewGosplanAPI(gpData, cfgGP, log, transcription)
	services := service.NewService(tgApi, gosplanApi, log, strCfg)
	handlers := middleware.NewHandler(services, log, strCfg, cfgHDL)
	servers := server.NewServer(handlers, log)

	servers.Run(serverCfg.Port)
}
