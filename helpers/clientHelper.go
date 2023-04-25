package helpers

import (
	"context"
	"time"

	configs "github.com/tarunrana0222/user_project_go/config"
	"github.com/tarunrana0222/user_project_go/models"
	"go.mongodb.org/mongo-driver/bson"
)

var clientCollection = configs.OpenCollection(configs.Client, "clients")

func ClientExists(clientId string) (models.Client, error) {
	var client models.Client

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	if err := clientCollection.FindOne(ctx, bson.M{"clientId": clientId}).Decode(&client); err != nil {
		return client, err
	}

	return client, nil
}
