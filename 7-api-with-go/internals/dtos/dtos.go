package dtos

type ResponseMeta struct {
	Page       int `json:"page,omitempty"`
	Limit      int `json:"limit,omitempty"`
	Total      int `json:"total,omitempty"`
	TotalPages int `json:"totalPages,omitempty"`
}
type APIResponse[T any] struct {
	Success bool         `json:"success"`
	Message string       `json:"message,omitempty"`
	Data    T            `json:"data,omitempty"`
	Meta    ResponseMeta `json:"meta,omitempty"`
}
