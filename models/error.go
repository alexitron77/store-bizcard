package models

type HTTPClientError struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

type HTTPBackendError struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"internal server error"`
}

type HTTPSuccess struct {
	Code    int    `json:"code" example:"200"`
	Message string `json:"message" example:"status success"`
}

type HTTPCreated struct {
	Code    int    `json:"code" example:"201"`
	Message string `json:"message" example:"status created"`
}
