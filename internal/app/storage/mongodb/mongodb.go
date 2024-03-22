package mongodb

import (
	"context"

	"github.com/markraiter/cardcheck/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	client *mongo.Client
	db     *mongo.Database
}

func New(ctx context.Context, cfg config.Mongo) (*Storage, error) {
	client, err := mongo.Connect(ctx,
		options.Client().
			ApplyURI(cfg.ConnectionString).SetAuth(
			options.Credential{
				Username: cfg.Username,
				Password: cfg.Password,
			},
		))
	if err != nil {
		return nil, err
	}

	db := client.Database(cfg.NameDB)

	return &Storage{client: client, db: db}, nil
}
