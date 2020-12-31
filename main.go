package main

import (
    "github.com/cemalkilic/jsonServer/controllers"
    "github.com/cemalkilic/jsonServer/database"
    "github.com/gin-gonic/gin"
    "github.com/go-playground/validator/v10"
)

func main() {
    router := gin.Default()

    database.Init()
    db := database.GetDB()

    v := validator.New()

    customEndpointController := controllers.NewCustomEndpointController(db, v)
    customEndpointController.SetDB(db)

    // Default handler to handle user routes
    router.NoRoute(customEndpointController.GetCustomEndpoint)
    router.POST("/addEndpoint", customEndpointController.AddCustomEndpoint)

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    router.Run() // serve on 0.0.0.0:8080
}
