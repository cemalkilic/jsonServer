package service

type GetEndpointParams struct{
    Endpoint string `json:"endpoint" binding:"required"`
}

type GetResponse struct {
    StatusCode int    `json:"statusCode"`
    Content    string `json:"content"`
    Err        error  `json:"err,omitempty"`
}

type AddEndpointParams struct {
    Username   string `json:"username"`
    Endpoint   string `json:"endpoint" binding:"required"`
    Content    string `json:"content" binding:"required"`
    StatusCode int    `json:"statusCode" binding:"required"`
}

type AddEndpointResponse struct {
    Endpoint string `json:"endpoint"`
    Err      error  `json:"err,omitempty"`
}
