package controllers

import (
    "github.com/cemalkilic/jsonServer/models"
    "github.com/cemalkilic/jsonServer/service"
    "github.com/gin-gonic/gin"
    "net/http"
)

type LoginController struct {
    loginService service.LoginService
    jWtService   service.JWTService
}

func NewLoginController(loginService service.LoginService, jwtService service.JWTService) *LoginController {
    return &LoginController{
        loginService: loginService,
        jWtService:   jwtService,
    }
}

func (controller *LoginController) Login(c *gin.Context) {
    var credential models.User
    err := c.ShouldBindJSON(&credential)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Login data must be a valid JSON!",
        })
        return
    }

    var token string
    isUserAuthenticated := controller.loginService.IsValidCredentials(credential)
    if isUserAuthenticated {
        token, err = controller.jWtService.GenerateToken(credential.Username)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "error": err.Error(),
            })
        }
    } else {
        c.JSON(http.StatusUnauthorized, gin.H{
            "error": "Given credentials did not match!",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "token": token,
    })
}

func (controller *LoginController) Signup(c *gin.Context) {
    var credential models.User
    err := c.ShouldBindJSON(&credential)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Signup data must be a valid JSON!",
        })
        return
    }

    err = controller.loginService.Signup(credential)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    // Return the token to user, to let her login right after signup
    token, err := controller.jWtService.GenerateToken(credential.Username)
    if err != nil {
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "success",
        "token": token,
    })
}
