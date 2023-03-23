package main

import (
	"log"
	"os"

	"Restaurant-Management/controllers"
	"Restaurant-Management/database"
	"Restaurant-Management/middleware"
	routes "Restaurant-Management/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	app := controllers.NewApplication(database.OpenCollection(database.Client, "food"), database.OpenCollection(database.Client, "user"))
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)
	routes.FoodRoutes(router)
	routes.MenuRoutes(router)

	router.Use(middleware.Authentication())
	router.GET("/addtocart", app.AddToCart())
	//router.DELETE("/removeitem", app.RemoveFoodItem())
	router.GET("/listcart", controllers.GetFoodFromCart())
	router.POST("/addaddress", controllers.AddAddress())
	router.GET("/cartcheckout", app.BuyFoodFromCart())
	router.GET("/instantbuy", app.InstantBuy())

	router.Run(":" + port)
}
