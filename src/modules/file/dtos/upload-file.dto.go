package dtos

import "mime/multipart"

type UploadFileDto struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}
