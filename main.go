package main

import (
	"net/http"

	"github.com/Gprisco/decanto-food-service/consul"
	"github.com/Gprisco/decanto-food-service/env"
	"github.com/Gprisco/decanto-food-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	consul.Register()
	r := gin.Default()
	baseUrl := env.GetInstance().BaseURL

	r.GET(baseUrl+"/check", (func(c *gin.Context) {
		c.JSON(http.StatusOK, nil)
	}))

	r.GET(baseUrl+"/recipe", handlers.GetRecipes)
	r.GET(baseUrl+"/recipe/:recipeId", handlers.GetRecipe)

	r.Run(env.GetInstance().Port)
}
