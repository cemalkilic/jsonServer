package database

import (
    "github.com/cemalkilic/jsonServer/models"
)

type DataStore interface {
    Insert(endpoint models.CustomEndpoint) error
    Select(username string, uri string) (models.CustomEndpoint, error)
    SelectByID(id int) (models.CustomEndpoint, error)
    SelectAllByUser(username string) ([]models.CustomEndpoint, error)
    Delete(id int) error
}

type UserStore interface {
    InsertUser(user models.User) error
    SelectByUsername(username string) (models.User, error)
}
