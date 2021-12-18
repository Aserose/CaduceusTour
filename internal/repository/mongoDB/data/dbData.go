package data

import "context"

func (g MongoData) Delete() {
	g.log.Info("gpData: deleting data")
	g.db.Drop(context.TODO())
}
