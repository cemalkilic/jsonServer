package main

import (
    "github.com/cemalkilic/jsonServer/controllers"
    "github.com/cemalkilic/jsonServer/database"
    "github.com/cemalkilic/jsonServer/utils/validator"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())

    router.Static("/", "./frontend/build")

    database.Init()
    db := database.GetDB()

    v := validator.NewValidator()

    customEndpointController := controllers.NewCustomEndpointController(db, v)
    customEndpointController.SetDB(db)

    // Default handler to handle user routes
    router.NoRoute(customEndpointController.GetCustomEndpoint)
    router.POST("/addEndpoint", customEndpointController.AddCustomEndpoint)

    router.Run() // serve on 0.0.0.0:8080
}
