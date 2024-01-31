package common

type Pagination struct {
	Total    int64 `json:"total,omitempty" form:"-"`
	Page     int   `json:"page,omitempty" form:"page,default=1" binding:"gte=1"`
	PageSize int   `json:"pageSize,omitempty" form:"page-size,default=10" binding:"gte=0,lte=1000"`
}

func NewPagination() *Pagination {
	return &Pagination{
		Total:    0,
		Page:     1,
		PageSize: 10,
	}
}
