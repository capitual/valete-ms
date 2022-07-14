package handlers

type Response struct {
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func errorResponse(err error) *Response {
	return &Response{
		Body:    nil,
		Success: false,
		Message: err.Error(),
	}
}

func successResponse(data interface{}) *Response {
	return &Response{
		Body:    data,
		Success: true,
	}
}
