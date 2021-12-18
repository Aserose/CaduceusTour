package gosplan

import (
	"github.com/Aserose/CaduceusTour/internal/config"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/Aserose/CaduceusTour/internal/repository/mongoDB/data"
	"github.com/Aserose/CaduceusTour/pkg/logger"
)

type GosplanApi interface {
	RequestToDataSource(params map[string]string) ([]models.ContractInfo, string)
}

type gosplan struct {
	gpData        *data.MongoData
	cfg           *config.GPConfig
	log           logger.Logger
	transcription config.TranscriptRespSource
}

func NewGosplanAPI(gpData *data.MongoData, cfg *config.GPConfig, log logger.Logger, transcription config.TranscriptRespSource) GosplanApi {
	return &gosplan{
		gpData:        gpData,
		cfg:           cfg,
		log:           log,
		transcription: transcription,
	}
}
