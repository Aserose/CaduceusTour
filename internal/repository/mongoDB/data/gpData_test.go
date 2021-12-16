package data

import (
	"context"
	"github.com/Aserose/CaduceusTour/internal/config"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/Aserose/CaduceusTour/internal/repository/mongoDB"
	"github.com/Aserose/CaduceusTour/pkg/logger"
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"os"
	"testing"
	"time"
)

const (
	testToken    = "32ktgk23kg[23k2pg32"
	updTestToken = "6436mgylom3gl34m"
)

func TestGPData(t *testing.T) {
	times, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	updTimes := times.AddDate(1, 0, 0)

	Convey("setup", t, func() {
		gpData, err := setupGpData()
		So(err, ShouldBeNil)

		Convey("put token", func() {
			putToken(gpData, times)

			Convey("get token", func() {
				getToken(gpData, times)
				So(os.Getenv("GP_TOKEN"), ShouldEqual, testToken)
				log.Print(getToken(gpData, times))

				Convey("overdue token", func() {
					So(getToken(gpData, updTimes), ShouldEqual, "overdue")
					log.Print(getToken(gpData, updTimes))

					Convey("update token", func() {
						updateToken(gpData, updTimes)

						Convey("get updated token", func() {
							getToken(gpData, updTimes)
							So(os.Getenv("GP_TOKEN"), ShouldEqual, updTestToken)
							log.Print(getToken(gpData, updTimes))

							Convey("drop", func() {
								gpData.Delete()
							})
						})
					})
				})
			})
		})
		Convey("get non-existent token", func() {
			So(getToken(gpData, times), ShouldEqual, "empty")
		})

	})
}

func TestGPOrganization(t *testing.T) {
	times, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	organization := models.Organization{
		"Pupkin Corporation",
		69,
		models.Payment{
			88005553535,
			9,
			300,
			[]int{300},
			777,
			666,
			"tractor purchase",
			"lollipop recycling",
			[]map[int]string{{777: "tractor purchase"}},
			[]map[int]string{{666: "lollipop recycling"}},
			[2][]int{{300}, {300}},
			models.Timeline{
				StartDate:   times.AddDate(0, 5, 0),
				EndDate:     times.AddDate(1, 0, 0),
				StartPeriod: times,
			}}}

	Convey("setup", t, func() {
		var listNames []string
		db, err := setupGpData()
		So(err, ShouldBeNil)

		Convey("put/get organization", func() {
			putOrganization(db, organization)
			So(getOrganization(db, organization.Name), ShouldResemble, organization)
			listNames = append(listNames, organization.Name)

			putOrganization(db, func() models.Organization { organization.Name = "Vasyan Industries"; return organization }())
			So(getOrganization(db, organization.Name), ShouldResemble, organization)
			listNames = append(listNames, organization.Name)

			Convey("get list names", func() {
				So(getNames(db), ShouldResemble, listNames)
				Convey("drop", func() {
					db.Delete()
				})
			})
		})
	})
}

func setupGpData() (*DBData, error) {
	_, cfgDB, _, _, _ := config.Init()
	log := logger.NewLogger()

	db, ctx, err := mongoDB.MongoInit(context.Background(), cfgDB.MongoDB.Host, cfgDB.MongoDB.Port,
		cfgDB.MongoDB.Username, cfgDB.MongoDB.Password, "testDatabase", log)
	if err != nil {
		return nil, err
	}

	return NewGpData(db, db.Collection("GPData"), db.Collection("AccessData"), ctx, log), nil
}

func updateToken(gpData *DBData, times time.Time) {
	gpData.UpdateToken(updTestToken, times)
}

func putToken(gpData *DBData, times time.Time) {
	gpData.PutToken(testToken, times)
}

func getToken(gpData *DBData, times time.Time) string {
	return gpData.GetToken(times)
}

func putOrganization(gpData *DBData, org models.Organization) {
	gpData.Put(org)
}

func getNames(gpData *DBData) []string {
	return gpData.GetListNames()
}

func getOrganization(gpData *DBData, name string) models.Organization {
	return gpData.Get(name)
}
