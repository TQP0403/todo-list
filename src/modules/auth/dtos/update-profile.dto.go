package dtos

type UpdateProfileDto struct {
	DisplayName string `json:"displayName" binding:"max=255,min=3"`
}
