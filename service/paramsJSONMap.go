package service

type GetEndpointParams struct{
    Endpoint string `json:"endpoint" validate:"required,uri"`
}

type GetResponse struct {
    StatusCode int    `json:"statusCode"`
    Content    string `json:"content"`
    Err        error  `json:"err,omitempty"`
}

type AddEndpointParams struct {
    Username   string `json:"username" validate:"omitempty,alphanum"`
    Endpoint   string `json:"endpoint" validate:"required,alphanum"`
    Content    string `json:"content" validate:"required"`
    StatusCode int    `json:"statusCode,string" validate:"required,numeric"`
}

type AddEndpointResponse struct {
    Endpoint string `json:"endpoint"`
    Err      error  `json:"err,omitempty"`
}
