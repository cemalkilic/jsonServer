package controllers

import (
    "github.com/cemalkilic/jsonServer/database"
    "github.com/cemalkilic/jsonServer/service"
    "github.com/cemalkilic/jsonServer/utils"
    "github.com/cemalkilic/jsonServer/utils/validator"
    "github.com/gin-gonic/gin"
    "strings"
)

type CustomEndpointController struct {
    dataStore database.DataStore
    validator *validator.CustomValidator
}

func NewCustomEndpointController(db database.DataStore, v *validator.CustomValidator) *CustomEndpointController {
    return &CustomEndpointController{
        dataStore: db,
        validator: v,
    }
}

func (cec *CustomEndpointController) SetDB(dataStore database.DataStore) {
    cec.dataStore = dataStore
}

func (cec *CustomEndpointController) AddCustomEndpoint(c *gin.Context) {
    var addEndpointRequest service.AddEndpointParams
    _ = c.ShouldBindJSON(&addEndpointRequest)

    srv := service.NewService(cec.dataStore, cec.validator)
    response, err := srv.AddEndpoint(addEndpointRequest)
    if err != nil {
        internalError(c, err)
        return
    }

    if e, ok := response.Err.(error); ok && e != nil {
        internalError(c, e)
        return
    }

    fullEndpointURL := utils.GetFullHTTPUrl(c.Request.Host, response.Endpoint, c.Request.TLS != nil)
    c.JSON(200, gin.H{
        "endpoint": fullEndpointURL,
    })
}

func (cec *CustomEndpointController) GetCustomEndpoint(c *gin.Context) {
    url := c.Request.URL.Path

    srv := service.NewService(cec.dataStore, cec.validator)
    response, err := srv.GetCustomEndpoint(service.GetEndpointParams{Endpoint: url})
    if err != nil {
        internalError(c, err)
        return
    }

    if e, ok := response.Err.(error); ok && e != nil {
        internalError(c, e)
        return
    }

    c.DataFromReader(response.StatusCode,
        int64(len(response.Content)),
        gin.MIMEJSON,
        strings.NewReader(response.Content), nil)
}
