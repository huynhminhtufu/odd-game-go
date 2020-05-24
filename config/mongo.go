package config

import (
	"context"
	"github.com/oddx-team/odd-game-server/pkg/l"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Config) NewMongoConnection(host, databaseName string) *mongo.Database {
	client, err := mongo.NewClient(options.Client().ApplyURI(host))
	if err != nil {
		ll.Panic("Mongo initial error", l.Error(err))
	}
	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		ll.Panic("Mongo connect error", l.Error(err))
	}

	ll.Info("Connected mongoDB at " + host + ", database " + databaseName)

	return client.Database(databaseName)
}
