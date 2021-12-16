package data

import (
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)

func (g DBData) Put(organization models.Organization) {
	log.Print("insert organization")
	_, err := g.gpData.InsertOne(g.ctx, organization)
	if err != nil {
		g.log.Errorf("gpData: %v", err.Error())
	}
}

func (g DBData) Get(name string) models.Organization {
	var organization models.Organization
	filter := bson.D{{"name", name}}

	err := g.gpData.FindOne(g.ctx, filter).Decode(&organization)
	if err != nil {
		g.log.Info(err.Error())
		g.log.Error(err.Error())
	}

	return organization
}

func (g DBData) GetListNames() []string {
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

func (g DBData) Update(organization models.Organization) {

}

func (g DBData) DeleteOrganization(name string) {
	filter := bson.D{{"name", name}}

	g.gpData.DeleteOne(g.ctx, filter)
}
