package domain

import "time"

const (
	DEFAULT_LIMIT  = int32(10)
	DEFAULT_OFFSET = int32(0)
)

type Response struct {
	Data interface{} `json:"data"`
}

type FetchResponse struct {
	Data interface{} `json:"data"`
	Next interface{} `json:"next"`
}

func NewResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}

func NewFetchResponse[T string | int32 | time.Time](data interface{}, next T) *FetchResponse {
	return &FetchResponse{
		Data: data,
		Next: next,
	}
}
