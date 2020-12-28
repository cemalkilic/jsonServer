package controllers

import (
    "github.com/cemalkilic/jsonServer/database"
    "github.com/cemalkilic/jsonServer/service"
    "github.com/gin-gonic/gin"
    "log"
    "strings"
)

type CustomEndpointController struct {
    dataStore database.DataStore
}

func (cec *CustomEndpointController) SetDB(dataStore database.DataStore) {
    cec.dataStore = dataStore
}

func (cec *CustomEndpointController) AddCustomEndpoint(c *gin.Context) {
    var addEndpointRequest service.AddEndpointParams
    if c.ShouldBindJSON(&addEndpointRequest) == nil {
        log.Println(addEndpointRequest)
    }

    srv := service.NewService(cec.dataStore)
    response, err := srv.AddEndpoint(addEndpointRequest)
    if err != nil {
        internalError(c, err)
        return
    }

    if e, ok := response.Err.(error); ok && e != nil {
        internalError(c, e)
        return
    }
    c.JSON(200, gin.H{
        "endpoint": response.Endpoint,
    })
}

func (cec *CustomEndpointController) GetCustomEndpoint(c *gin.Context) {
    url := c.Request.URL.Path

    srv := service.NewService(cec.dataStore)
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
