package service

import (
    "errors"
    "fmt"
    "github.com/cemalkilic/jsonServer/database"
    "github.com/cemalkilic/jsonServer/models"
    "github.com/cemalkilic/jsonServer/utils/validator"
    "strings"
)

type jsonService struct {
    db database.DataStore
    validate *validator.CustomValidator
}

func NewService(db database.DataStore, v *validator.CustomValidator) *jsonService {
    return &jsonService{
        db: db,
        validate: v,
    }
}

func (srv *jsonService) GetCustomEndpoint(params GetEndpointParams) (GetResponse, error) {

    // Terminate the request if the input is not valid
    if err := srv.validate.ValidateStruct(params); err != nil {
        return GetResponse{}, err
    }

    endpoint := strings.Trim(params.Endpoint, "/")
    if len(endpoint) == 0 {
        return GetResponse{}, errors.New("empty URI given")
    }

    urlParts := strings.Split(endpoint, "/")
    username := urlParts[0]
    if len(urlParts[1:]) < 1 {
      return GetResponse{}, errors.New("URI given without custom endpoint")
    }

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
    // Prepend with a slash to behave it like a uri
    if !strings.HasPrefix(params.Endpoint, "/") {
        params.Endpoint = "/" + params.Endpoint
    }

    // Terminate the request if the input is not valid
    if err := srv.validate.ValidateStruct(params); err != nil {
       return AddEndpointResponse{}, err
    }

    // Trim the slashes after validation :/ That's way easier than custom validation
    params.Endpoint = strings.Trim(params.Endpoint, "/")

    // Use the default username if not exists in the params
    username := params.Username
    if username == "" {
        username = "guest"
    }

    // Make sure the same endpoint does not already exist
    response, err := srv.db.Select(username, params.Endpoint)
    if err != nil {
        return AddEndpointResponse{}, err
    }

    if response.ID != 0 {
        return AddEndpointResponse{}, errors.New("endpoint already exists")
    }

    endpointObj := models.CustomEndpoint{
        Username:   username,
        URI:        params.Endpoint,
        Content:    params.Content,
        StatusCode: params.StatusCode,
    }

    err = srv.db.Insert(endpointObj)
    if err != nil {
        return AddEndpointResponse{}, err
    }

    // Complete path of created endpoint
    userEndpoint := fmt.Sprintf("%s/%s", username, params.Endpoint)

    return AddEndpointResponse{
        Endpoint: userEndpoint,
        Err:      nil,
    }, nil
}
