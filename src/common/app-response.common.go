package common

type SuccessResponse struct {
	Message  string      `json:"message,omitempty"`
	Metadata *Pagination `json:"metadata,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func NewSimpleResponse() SuccessResponse {
	return SuccessResponse{Message: "OK"}
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{Message: "OK", Data: data}
}

func NewListResponse(data PaginationResponse) SuccessResponse {
	return SuccessResponse{Message: "OK", Data: data.Rows, Metadata: &data.Metadata}
}

func NewErrorResponse(err AppError) ErrorResponse {
	return ErrorResponse{Message: "Fail", Error: err.Message}
}
