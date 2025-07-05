package mongodb

import (
	"authserver/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoDB *mongo.Client

// Call this once during startup
func InitMongo() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://172.20.78.91:27017"))
	if err != nil {
		return err
	}

	MongoDB = client
	return nil
}

func InsertAuditLog(log models.AuditLog) error {
	collection := MongoDB.Database("authDB").Collection("logs")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, log)
	return err
}
