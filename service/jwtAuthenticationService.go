package service

import (
    "errors"
    "os"
    "time"

    "github.com/dgrijalva/jwt-go"
)

type JWTService interface {
    GenerateToken(email string) (string, error)
    ValidateToken(token string) (*jwt.Token, error)
}

type authClaims struct {
    Username string `json:"username"`
    jwt.StandardClaims
}

type jwtServices struct {
    secretKey string
    issuer    string
}

func JWTAuthService() JWTService {
    return &jwtServices{
        secretKey: getSecretKey(),
        issuer:    "CK-JSONServer",
    }
}

func getSecretKey() string {
    secret := os.Getenv("SECRET")
    if secret == "" {
        secret = "mysecret"
    }
    return secret
}

func (service *jwtServices) GenerateToken(username string) (string, error) {
    claims := &authClaims{
        username,
        jwt.StandardClaims{
            Issuer:    service.issuer,
            IssuedAt:  time.Now().Unix(),
            ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    t, err := token.SignedString([]byte(service.secretKey))
    if err != nil {
        return "", err
    }

    return t, nil
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
    return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
        // Make sure given alg is the one we set
        if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
            return nil, errors.New("invalid token alg")
        }
        return []byte(service.secretKey), nil
    })
}
