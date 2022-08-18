package services

import (
	"context"

	"github.com/Gprisco/decanto-food-service/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetRecipes(page int64, limit int64) []models.Recipe {
	findOptions := options.Find()

	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)

	var recipes []models.Recipe
	cursor, err := models.RecipeCollection.Find(context.TODO(), bson.D{}, findOptions)

	if err != nil {
		panic(err)
	}

	cursor.All(context.TODO(), &recipes)
	return recipes
}

func GetRecipe(id primitive.ObjectID) models.Recipe {
	var recipe models.Recipe
	err := models.RecipeCollection.FindOne(context.TODO(), bson.D{{Key: "_id", Value: id}}).Decode(&recipe)

	if err != nil {
		panic(err)
	}

	return recipe
}
