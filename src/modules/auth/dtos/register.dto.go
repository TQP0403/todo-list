package dtos

type RegisterDto struct {
	DisplayName string `json:"displayName" binding:"required"`
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
}
