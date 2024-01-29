package dtos

type LoginDto struct {
	Username string `json:"username" binding:"required,max=255,min=3"`
	Password string `json:"password" binding:"required,max=255,min=3"`
}
