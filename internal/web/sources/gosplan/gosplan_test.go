package gosplan

import (
	"context"
	"github.com/Aserose/CaduceusTour/internal/config"
	"github.com/Aserose/CaduceusTour/internal/repository/mongoDB"
	"github.com/Aserose/CaduceusTour/internal/repository/mongoDB/data"
	"github.com/Aserose/CaduceusTour/pkg/logger"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGosplan(t *testing.T) {

	Convey("setup", t, func() {
		gosplan, err := setup()
		if err != nil {
			err.Error()
		}

		Convey("request", func() {
			a, _ := gosplan.RequestToDataSource(map[string]string{"name": "Эрмитаж", "time": time.Now().Format("2006-01-02T15:04:05.000Z")})
			So(a[1], ShouldNotBeNil)
		})
	})

}

func setup() (GosplanApi, error) {
	logs := logger.NewLogger()
	_, cfgDB, cfgGP, _, err := config.Init()
	if err != nil {
		logs.Error(err.Error())
	}

	_, transcription,_ := config.InitStrCfg("configs/configs.yml")

	db, _, err := mongoDB.MongoInit(context.Background(), cfgDB.MongoDB.Host, cfgDB.MongoDB.Port,
		cfgDB.MongoDB.Username, cfgDB.MongoDB.Password, cfgDB.MongoDB.DBName, logs)
	if err != nil {

	}

	gpData := data.NewMongoData(db, db.Collection(cfgDB.MongoDB.GPData), db.Collection(cfgDB.MongoDB.Access), context.Background(), logs)

	return NewGosplanAPI(gpData, cfgGP, logs, transcription), nil
}
