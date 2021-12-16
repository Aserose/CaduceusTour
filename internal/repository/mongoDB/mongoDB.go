package mongoDB

import (
	"context"
	"fmt"
	"github.com/Aserose/CaduceusTour/pkg/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoInit(ctx context.Context, host, port, username, password, database string, log logger.Logger) (*mongo.Database, context.Context, error) {
	log.Info("repository: initialization")

	var mongoDBURL string

	var isAuth bool

	if username == "" && password == "" {
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s", host, port)
	} else {
		isAuth = true
		mongoDBURL = fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	}

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	if isAuth {
		clientOptions.SetAuth(options.Credential{
			Username: username,
			Password: password,
		})
	}

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to mongoDB due to error: %v", err.Error())
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, nil, fmt.Errorf("failed to ping mongoDB due to error: %v", err.Error())
	}

	log.Info("repository: initialization ok")

	return client.Database(database), ctx, nil
}
