package common

type SuccessResponse struct {
	Message  string      `json:"message,omitempty"`
	Metadata *Pagination `json:"metadata,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Log     string `json:"log,omitempty"`
}

func NewSimpleResponse() SuccessResponse {
	return SuccessResponse{Message: "ok"}
}

func NewSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{Message: "ok", Data: data}
}

func NewListResponse(data PaginationResponse) SuccessResponse {
	return SuccessResponse{Message: "ok", Data: data.Rows, Metadata: &data.Metadata}
}

func NewErrorResponse(err AppError) ErrorResponse {
	return ErrorResponse{Message: "error", Error: err.Message, Log: err.LogError}
}
