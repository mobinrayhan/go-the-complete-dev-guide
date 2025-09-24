package dtos

type ResponseMeta struct {
	Total      int  `json:"total,omitempty"`
	PerPage    int  `json:"perPage,omitempty"`
	Page       int  `json:"page,omitempty"`
	TotalPages int  `json:"totalPages,omitempty"`
	HasNext    bool `json:"hasNext"`
	HasPrev    bool `json:"hasPrev"`
}
type ApiResponseList[T any] struct {
	Success bool          `json:"success,omitempty"`
	Data    T             `json:"data"`
	Message string        `json:"message,omitempty"`
	Meta    *ResponseMeta `json:"meta,omitempty"`
}

type ApiResponseSingle[T any] struct {
	Success bool   `json:"success"`
	Data    *T     `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}
