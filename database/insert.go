package database

import (
	"context"
)

func Insert(collection string, data interface{})  error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("database-1").Collection(collection)
	_, err := c.InsertOne(context.Background(), data)

	return err
}