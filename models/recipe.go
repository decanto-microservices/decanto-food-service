package models

import (
	"github.com/Gprisco/decanto-food-service/db"
	"github.com/Gprisco/decanto-food-service/env"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Recipe struct {
	ID   primitive.ObjectID `bson:"_id" json:"_id"`
	Name string             `bson:"name" json:"name"`
}

var RecipeCollection *mongo.Collection = db.GetInstance().Client().Database(env.GetInstance().DB).Collection("f_recipe")
