package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(username string) ([]byte, error) {
	client, ctx := getConnection()

	defer client.Disconnect(ctx)
	c := client.Database(databaseName).Collection("users")
	opts := options.Find()
	curr, err := c.Find(ctx, bson.D{{Key: "username", Value: username}}, opts)

	if err != nil {
		return nil, err
	} else {
		var elem []byte
		err := curr.Decode(elem)
		if err != nil {
			fmt.Printf("Erro ao buscar: %v", err)
			return nil, fmt.Errorf("user_not_found")
		}
		return elem, nil
	}
}

func Get(collection string) (interface{}, error) {
	client, ctx := getConnection()

	defer client.Disconnect(ctx)

	c := client.Database("websocket").Collection(collection)
	opts := options.Find().SetLimit(6)
	curr, err := c.Find(
		context.TODO(),
		bson.D{},
		opts,
	)

	if err != nil {
		return nil, err
	}

	var results []interface{}

	for curr.Next(context.TODO()) {
		var elem interface{}
		err := curr.Decode(&elem)
		if err != nil {
			fmt.Printf("Erro ao buscar: %v", err)
		}

		results = append(results, elem)
	}
	return results, nil

}
