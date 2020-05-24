package stores

import (
	"context"

	"github.com/oddx-team/odd-game-server/config"
	"github.com/oddx-team/odd-game-server/pkg/l"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	ll = l.New()
)

type MainStore struct {
	*mongo.Database
}

func NewMongoConnection(cfg *config.Config) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(cfg.Mongo.Host))
	if err != nil {
		ll.Panic("Mongo initial error", l.Error(err))
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		ll.Panic("Mongo connect error", l.Error(err))
	}

	ll.Info("Connected mongoDB at " + cfg.Mongo.Host + ", database " + cfg.Mongo.DatabaseName)

	return client.Database(cfg.Mongo.DatabaseName)
}
