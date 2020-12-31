package service

type GetEndpointParams struct{
    Endpoint string `json:"endpoint" validate:"required"`
}

type GetResponse struct {
    StatusCode int    `json:"statusCode"`
    Content    string `json:"content"`
    Err        error  `json:"err,omitempty"`
}

type AddEndpointParams struct {
    Username   string `json:"username"`
    Endpoint   string `json:"endpoint" validate:"required"`
    Content    string `json:"content" validate:"required"`
    StatusCode int    `json:"statusCode,string" validate:"required"`
}

type AddEndpointResponse struct {
    Endpoint string `json:"endpoint"`
    Err      error  `json:"err,omitempty"`
}
