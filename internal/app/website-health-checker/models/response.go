package models

type JSONResponse struct {
	Meta JSONResponseMeta `json:"meta"`
	Data interface{}      `json:"data,omitempty"`
}

type JSONResponseMeta struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
