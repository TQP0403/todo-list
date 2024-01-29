package common

type Pagination struct {
	Total    int64 `json:"total,omitempty" form:"-"`
	Page     int   `json:"page,omitempty" form:"page,default=1" binding:"gte=1"`
	PageSize int   `json:"pageSize,omitempty" form:"page-size,default=10" binding:"gte=0,lte=1000"`
}

type PaginationResponse struct {
	Metadata *Pagination `json:"metadata,omitempty"`
	Rows     interface{} `json:"rows,omitempty"`
}

func NewPagination() *Pagination {
	return &Pagination{
		Total:    0,
		Page:     1,
		PageSize: 10,
	}
}

func NewPaginationResponse(metadata *Pagination, rows interface{}) *PaginationResponse {
	return &PaginationResponse{
		Metadata: metadata,
		Rows:     rows,
	}
}

func (data *PaginationResponse) GetSuccessResponse() *SuccessResponse {
	return &SuccessResponse{Message: "ok", Data: data.Rows, Metadata: data.Metadata}
}
