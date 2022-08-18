package handlers

import (
	"net/http"
	"strconv"

	"github.com/Gprisco/decanto-food-service/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRecipes(c *gin.Context) {
	page, err := strconv.ParseInt(c.Query("page"), 10, 32)
	limit, err := strconv.ParseInt(c.Query("limit"), 10, 32)

	if err != nil || page <= 0 || limit <= 0 {
		c.JSON(http.StatusBadRequest, "'page' and 'limit' are required and should be > 0")
		return
	}

	recipes := services.GetRecipes(page, limit)
	c.JSON(http.StatusOK, recipes)
}

func GetRecipe(c *gin.Context) {
	recipeId, err := primitive.ObjectIDFromHex(c.Param("recipeId"))

	if err != nil {
		panic(err)
	}

	recipe := services.GetRecipe(recipeId)
	c.JSON(http.StatusOK, recipe)
}
