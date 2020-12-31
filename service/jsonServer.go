package service

import (
    "errors"
    "github.com/cemalkilic/jsonServer/database"
    "github.com/cemalkilic/jsonServer/models"
    "github.com/cemalkilic/jsonServer/utils"
    "github.com/go-playground/validator/v10"
    "strings"
)

type jsonService struct {
    db database.DataStore
    validate *validator.Validate
}

func NewService(db database.DataStore, v *validator.Validate) *jsonService {
    return &jsonService{
        db: db,
        validate: v,
    }
}

func (srv *jsonService) GetCustomEndpoint(params GetEndpointParams) (GetResponse, error) {

    // Terminate the request if the input is not valid
    if err := srv.validate.Struct(params); err != nil {
        return GetResponse{}, err
    }

    endpoint := strings.Trim(params.Endpoint, "/")
    urlParts := strings.Split(endpoint, "/")
    username := urlParts[0]
    userEndpoint := strings.Join(urlParts[1:], "/")

    customEndpoint, err := srv.db.Select(username, userEndpoint)
    if err != nil {
        return GetResponse{}, err
    }

    if customEndpoint.ID == 0 {
        // not found the custom endpoint
        return GetResponse{}, errors.New("404: Not Found")
    }

    return GetResponse{
        StatusCode: customEndpoint.StatusCode,
        Content:    customEndpoint.Content,
        Err:        nil,
    }, nil
}

func (srv *jsonService) AddEndpoint(params AddEndpointParams) (AddEndpointResponse, error) {
    // Terminate the request if the input is not valid
    if err := srv.validate.Struct(params); err != nil {
       return AddEndpointResponse{}, err
    }

    // Create a random username if not exists in the params
    username := params.Username
    if username == "" {
        username = utils.GetRandomUsername()
    }

    endpointObj := models.CustomEndpoint{
        Username:   username,
        URI:        params.Endpoint,
        Content:    params.Content,
        StatusCode: params.StatusCode,
    }

    err := srv.db.Insert(endpointObj)
    if err != nil {
        return AddEndpointResponse{}, err
    }

    return AddEndpointResponse{
        Endpoint: params.Endpoint,
        Err:      nil,
    }, nil
}
