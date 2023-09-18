package database

import (
	"log"
)

func Insert(collection string, data interface{})  error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(databaseName).Collection(collection)
	result, err := c.InsertOne(ctx, data)


	log.Println(result)

	return err
}