package main

import (
    "github.com/cemalkilic/jsonServer/controllers"
    "github.com/cemalkilic/jsonServer/database"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"

    "github.com/gin-contrib/cors"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())

    router.Static("/", "./frontend/build")

    database.Init()
    db := database.GetDB()

    v := validator.New()

    customEndpointController := controllers.NewCustomEndpointController(db, v)
    customEndpointController.SetDB(db)

    // Default handler to handle user routes
    router.NoRoute(customEndpointController.GetCustomEndpoint)
    router.POST("/addEndpoint", customEndpointController.AddCustomEndpoint)

    router.Run() // serve on 0.0.0.0:8080
}
