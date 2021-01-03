package main

import (
    "github.com/cemalkilic/jsonServer/config"
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

    router.StaticFile("/", "./frontend/build/index.html")
    router.Static("/static", "./frontend/build/static")

    cfg, _ := config.LoadConfig(".")

    mysqlHandler := database.NewMySQLDBHandler(cfg)
    dataStore := database.GetSQLDataStore(mysqlHandler)
    userStore := database.GetSQLUserStore(mysqlHandler)

    v := validator.NewValidator()

    customEndpointController := controllers.NewCustomEndpointController(dataStore, v)
    customEndpointController.SetDB(dataStore)

    loginService := service.DBLoginService(userStore, v)
    jwtService := service.JWTAuthService(cfg)
    loginController := controllers.NewLoginController(loginService, jwtService)

    router.POST("/login", loginController.Login)
    router.POST("/signup", loginController.Signup)
    router.GET("/user/me", middlewares.AuthorizeJWT(jwtService), func(context *gin.Context) {
        context.JSON(200, gin.H{
            "success": true,
        })
    })

    // Default handler to handle user routes
    router.NoRoute(customEndpointController.GetCustomEndpoint)
    router.POST("/addEndpoint", middlewares.AuthorizeJWT(jwtService), customEndpointController.AddCustomEndpoint)

    router.Run(cfg.ServerAddress)
}
