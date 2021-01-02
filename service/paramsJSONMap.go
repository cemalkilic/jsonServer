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
    Endpoint   string `json:"endpoint" validate:"required,uri"`
    Content    string `json:"content" validate:"required"`
    StatusCode int    `json:"statusCode" validate:"required,httpStatus"`
}

type AddEndpointResponse struct {
    Endpoint string `json:"endpoint"`
    Err      error  `json:"err,omitempty"`
}
