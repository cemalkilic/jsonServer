package service

import "github.com/cemalkilic/jsonServer/models"

type LoginService interface {
    IsValidCredentials(user models.User) bool
}

type loginInformation struct {
    username string
    password string
}

func StaticLoginService() LoginService {
    return &loginInformation{
        username: "testing",
        password: "pass",
    }
}

func (info *loginInformation) IsValidCredentials(user models.User) bool {
    return info.username == user.Username && info.password == user.Password
}
