package dtos

type ResponseMeta struct {
	Total    int `json:"total,omitempty"`
	Page     int `json:"page,omitempty"`
	PerPages int `json:"perPage,omitempty"`
	Limit    int `json:"limit,omitempty"`
}
type ApiResponseList[T any] struct {
	Success bool          `json:"success,omitempty"`
	Data    T             `json:"data,omitempty"`
	Message string        `json:"message,omitempty"`
	Meta    *ResponseMeta `json:"meta,omitempty"`
}

type ApiResponseSingle[T any] struct {
	Success bool   `json:"success,omitempty"`
	Data    *T     `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
