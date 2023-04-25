package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	configs "github.com/tarunrana0222/user_project_go/config"
	"github.com/tarunrana0222/user_project_go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userCollection = configs.OpenCollection(configs.Client, "users")
	validate       = validator.New()
)

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user.Client = c.GetString("clientId")
		user.ID = primitive.NewObjectID()
		if validation := validate.Struct(&user); validation != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validation.Error()})
			return
		}

		insertedId, err := userCollection.InsertOne(ctx, user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"data": insertedId})
	}
}

func UpdateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		var updatedUser models.User
		user.Client = c.GetString("clientId")
		if err := userCollection.FindOneAndUpdate(ctx,
			bson.D{{"client", user.Client}, {"userId", user.UserID}},
			bson.M{"$set": bson.M{"os": user.Os, "name": user.Name}},
			options.FindOneAndUpdate().SetReturnDocument(options.After),
		).Decode(&updatedUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": updatedUser})
	}
}

func GetAllUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		client := c.GetString("clientId")
		users := []models.User{}

		cursor, err := userCollection.Find(ctx, bson.M{"client": client})
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		for cursor.Next(ctx) {
			var user models.User
			if err = cursor.Decode(&user); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			users = append(users, user)
		}

		c.JSON(200, gin.H{"data": users})
	}
}

func GetSingleUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userId missing"})
			return
		}
		client := c.GetString("clientId")

		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := userCollection.FindOne(ctx, bson.D{{"client", client}, {"userId", userId}}).Decode(&user); err != nil {
			if err.Error() == "mongo: no documents in result" {
				c.JSON(200, gin.H{"data": nil, "message": "User not found"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user})
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("userId")
		if userId == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "userId missing"})
			return
		}
		client := c.GetString("clientId")
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		if err := userCollection.FindOneAndDelete(ctx, bson.D{{"client", client}, {"userId", userId}}).Decode(&user); err != nil {
			if err.Error() == "mongo: no documents in result" {
				c.JSON(200, gin.H{"data": user, "message": "User not found to be deleted"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"data": user, "message": "User Deleted"})
	}
}
