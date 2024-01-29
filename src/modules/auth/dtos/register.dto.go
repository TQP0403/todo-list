package dtos

type RegisterDto struct {
	DisplayName string `json:"displayName" binding:"required,max=255,min=3"`
	Username    string `json:"username" binding:"required,max=255,min=3"`
	Password    string `json:"password" binding:"required,max=255,min=3"`
}
