package middlewares

import (
    "fmt"
    "github.com/cemalkilic/jsonServer/service"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc{
    return func(c *gin.Context) {
        const BearerSchema = "Bearer"

        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            return
            /*
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
                "error": "Authorization token is not given!",
            })
            return
            */
        }

        if !strings.HasPrefix(authHeader, BearerSchema) {
            return
            /*
            c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
                "error": "Authorization token must be type of Bearer!",
            })
            return
            */
        }

        tokenString := authHeader[len(BearerSchema):]
        tokenString = strings.TrimSpace(tokenString)

        token, err := jwtService.ValidateToken(tokenString)
        if err != nil {
            fmt.Printf("%v", err)
            c.AbortWithStatusJSON(400, gin.H{
                "error": "Given token is invalid!",
            })
            return
        }

        if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
            c.Set("username", claims["username"])
            //c.Next()
        } else {
            c.JSON(http.StatusUnauthorized, gin.H{
                "error": "Given JWT is invalid!",
            })
        }
    }
}
