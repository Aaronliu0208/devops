package injector

import (
	"context"
	"time"

	"casicloud.com/ylops/backend/internal/app/config"
	imongo "casicloud.com/ylops/backend/internal/app/model/impl/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

// InitMongo 初始化mongo存储
func InitMongo() (*mongo.Client, func(), error) {
	cfg := config.C.Mongo
	client, cleanFunc, err := imongo.NewClient(&imongo.Config{
		URI:      cfg.URI,
		Database: cfg.Database,
		Timeout:  time.Duration(cfg.Timeout) * time.Second,
	})
	if err != nil {
		return nil, cleanFunc, err
	}

	err = imongo.CreateIndexes(context.Background(), client)
	if err != nil {
		return nil, cleanFunc, err
	}

	return client, cleanFunc, nil
}
