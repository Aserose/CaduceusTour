package data

import (
	"context"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/Aserose/CaduceusTour/pkg/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoGPOrganizationData struct {
	gpData *mongo.Collection
	ctx context.Context
	log logger.Logger
}

func NewGPOrganizationData(gpData *mongo.Collection, ctx context.Context, log logger.Logger) *MongoGPOrganizationData {
	return &MongoGPOrganizationData{
		gpData: gpData,
		ctx: ctx,
	log: log,
	}
}

func (g MongoGPOrganizationData) Put(organization models.Organization) {
	log.Print("insert organization")
	_, err := g.gpData.InsertOne(g.ctx, organization)
	if err != nil {
		g.log.Errorf("gpData: %v", err.Error())
	}
}

func (g MongoGPOrganizationData) Get(name string) models.Organization {
	var organization models.Organization
	filter := bson.D{{"name", name}}

	err := g.gpData.FindOne(g.ctx, filter).Decode(&organization)
	if err != nil {
		g.log.Info(err.Error())
		g.log.Error(err.Error())
	}

	return organization
}

func (g MongoGPOrganizationData) GetListNames() []string {
	filter := bson.D{}
	var listNames []string

	cur, _ := g.gpData.Find(g.ctx, filter)

	for cur.Next(g.ctx) {
		var organization models.Organization
		err := cur.Decode(&organization)
		if err != nil {
			log.Fatal(err)
		}
		listNames = append(listNames, organization.Name)
	}
	return listNames
}

func (g MongoGPOrganizationData) Update(organization models.Organization) {

}

func (g MongoGPOrganizationData) DeleteOrganization(name string) {
	filter := bson.D{{"name", name}}

	g.gpData.DeleteOne(g.ctx, filter)
}
