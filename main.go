package main

import (
    "github.com/cemalkilic/jsonServer/controllers"
    "github.com/cemalkilic/jsonServer/database"
    "github.com/cemalkilic/jsonServer/middlewares"
    "github.com/cemalkilic/jsonServer/service"
    "github.com/cemalkilic/jsonServer/utils/validator"
    "github.com/gin-contrib/cors"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())

    router.Static("/", "./frontend/build")

    mysqlHandler := database.NewMySQLDBHandler()
    dataStore := database.GetSQLDataStore(mysqlHandler)
    userStore := database.GetSQLUserStore(mysqlHandler)

    v := validator.NewValidator()

    customEndpointController := controllers.NewCustomEndpointController(dataStore, v)
    customEndpointController.SetDB(dataStore)

    loginService := service.DBLoginService(userStore, v)
    jwtService := service.JWTAuthService()
    loginController := controllers.NewLoginController(loginService, jwtService)

    router.POST("/login", loginController.Login)
    router.POST("/signup", loginController.Signup)

    // Default handler to handle user routes
    router.NoRoute(customEndpointController.GetCustomEndpoint)
    router.POST("/addEndpoint", middlewares.AuthorizeJWT, customEndpointController.AddCustomEndpoint)

    router.Run() // serve on 0.0.0.0:8080
}
