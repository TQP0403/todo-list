package dtos

type ChangePasswordDto struct {
	OldPassword string `json:"oldPassword" binding:"required,max=255,min=3"`
	NewPassword string `json:"newPassword" binding:"required,max=255,min=3"`
}
