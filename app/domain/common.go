package domain

const (
	DEFAULT_LIMIT  = int32(10)
	DEFAULT_OFFSET = int32(0)
)

type Response struct {
	Data interface{} `json:"data"`
}

func NewResponse(data interface{}) *Response {
	return &Response{
		Data: data,
	}
}
