package model

type JSONResponse struct {
	Success bool  `json:"success"`
	Data    any   `json:"data,omitempty"`
	Error   error `json:"error,omitempty"`
}

func NewSuccessResponse(data any) *JSONResponse {
	return &JSONResponse{
		Success: true,
		Data:    data,
	}
}

func NewErrorResponse(err error) *JSONResponse {
	return &JSONResponse{
		Success: false,
		Error:   err,
	}
}
