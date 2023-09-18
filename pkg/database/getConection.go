package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getConnection() (client *mongo.Client, ctx context.Context) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:root@db:27017/goplaningdb?authSource=admin"))

	if err != nil {
		log.Printf("Não foi possivel estabelecer a conexão com o banco: %v", err)
		panic(err)
	}
	return
}
