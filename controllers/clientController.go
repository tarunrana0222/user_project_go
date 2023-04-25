package controllers

import (
	"context"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	configs "github.com/tarunrana0222/user_project_go/config"
	"github.com/tarunrana0222/user_project_go/helpers"
	"github.com/tarunrana0222/user_project_go/models"
	"go.mongodb.org/mongo-driver/bson"
)

var clientCollection = configs.OpenCollection(configs.Client, "clients")

type SignedDetails struct {
	Client string `json:"client"`
	jwt.StandardClaims
}

func GetAllClients() gin.HandlerFunc {
	return func(c *gin.Context) {
		clients := []models.Client{}
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		cursor, err := clientCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// for cursor.Next(ctx) {
		// 	var client models.Client
		// 	if err = cursor.Decode(&client); err != nil {
		// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		// 	}

		// 	clients = append(clients, client)
		// }

		defer cursor.Close(ctx)
		if err = cursor.All(context.TODO(), &clients); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"success": true, "data": clients})
	}
}

func GetAuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		client := c.GetHeader("x-client-key")
		if client == "" {
			c.JSON(400, gin.H{"message": "x-client-key header missing"})
			return
		}
		clientObj, err := helpers.ClientExists(client)
		if err != nil {
			c.JSON(400, gin.H{"message": "Invalid Client", "error": err.Error()})
			return
		}
		token, err := helpers.GenerateJwtToken(clientObj.ClientID)
		if err != nil {
			c.JSON(500, gin.H{"message": "Internal Server error", "error": err.Error()})
			return

		}

		c.JSON(200, gin.H{"token": token})
	}
}
