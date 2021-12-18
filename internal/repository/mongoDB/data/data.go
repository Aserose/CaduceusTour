package data

import (
	"context"
	"github.com/Aserose/CaduceusTour/internal/repository/models"
	"github.com/Aserose/CaduceusTour/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Data interface {
	Delete()
}

type AccessData interface {
	PutToken(token string, time time.Time)
	GetToken(times time.Time) string
	UpdateToken(token string, time time.Time)
}

type OrganizationData interface {
	Put(organization models.Organization)
	Get(name string) models.Organization
	Update(organization models.Organization)
	GetListNames() []string
	DeleteOrganization(name string)
}

type MongoData struct {
	db         *mongo.Database
	gpData     *mongo.Collection
	accessData *mongo.Collection
	ctx        context.Context
	log        logger.Logger
	Data
	AccessData
	OrganizationData
}

func NewMongoData(db *mongo.Database, gpData *mongo.Collection, accessData *mongo.Collection, ctx context.Context, log logger.Logger) *MongoData {
	return &MongoData{
		db:         db,
		ctx:        ctx,
		log:        log,
		gpData:     gpData,
		accessData: accessData,
		AccessData: NewGPData(accessData,ctx,log),
		OrganizationData: NewGPOrganizationData(gpData,ctx,log),
	}
}
