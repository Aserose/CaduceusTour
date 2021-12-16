package data

import (
	"context"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"time"
)

func (g DBData) PutToken(token string, time time.Time) {
	g.log.Info("gpData: put token")
	gpScheme := models.GP{
		Token:    token,
		TokenTTL: time,
	}

	_, err := g.accessData.InsertOne(g.ctx, gpScheme)
	if err != nil {
		g.log.Errorf("gpData: %v", err.Error())
	}

	toEnv(token)
}

func (g DBData) UpdateToken(token string, time time.Time) {
	g.log.Info("gpData: update token")
	filter := bson.D{}
	update := bson.D{{"$set", bson.D{{"token", token}, {"tokenttl", time}}}}

	_, err := g.accessData.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		g.log.Errorf("gpData: %v", err.Error())
	}

	toEnv(token)
}

func (g DBData) GetToken(times time.Time) string {
	g.log.Info("gpData: get token")
	var gpScheme models.GP
	filter := bson.D{}

	err := g.accessData.FindOne(g.ctx, filter).Decode(&gpScheme)
	if err != nil {
		g.log.Info(err.Error())
		return "empty"
	}

	if times.Unix() >= gpScheme.TokenTTL.AddDate(0, 0, 1).Unix() {
		return "overdue"
	}

	toEnv(gpScheme.Token)

	return "ok"
}

func (g DBData) Delete() {
	g.log.Info("gpData: deleting data")
	g.db.Drop(context.TODO())
}

func toEnv(token string) {
	if err := os.Setenv("GP_TOKEN", token); err != nil {
		log.Print(err.Error())
	}
}
