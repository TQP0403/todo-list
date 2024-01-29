package common

type SuccessResponse struct {
	Message  string      `json:"message"`
	Metadata *Pagination `json:"metadata,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

func NewSimpleResponse() *SuccessResponse {
	return &SuccessResponse{Message: "ok"}
}

func NewSuccessResponse(data interface{}) *SuccessResponse {
	return &SuccessResponse{Message: "ok", Data: data}
}
