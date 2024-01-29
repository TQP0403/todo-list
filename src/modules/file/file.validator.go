package file

import (
	"TQP0403/todo-list/src/helper"
	"errors"
	"mime/multipart"
	"os"
	"slices"
)

type IFileValidator interface {
	Validate(file *multipart.FileHeader) error
}

type FileValidator struct {
	allowTypes []string
	fileLimit  int64
}

func NewValidator() *FileValidator {
	return &FileValidator{
		allowTypes: []string{"image/jpeg", "image/png", "image/webp"},
		fileLimit: helper.GetDefaultNumber(
			int64(helper.ParseFloat(os.Getenv("UPLOAD_FILE_LIMIT"))*1024*1024),
			5*1024*1024, // 5mb
		),
	}
}

func (validator *FileValidator) Validate(file *multipart.FileHeader) error {
	if file.Size > validator.fileLimit {
		return errors.New("file too large")
	}

	if !slices.Contains(validator.allowTypes, file.Header.Get("Content-Type")) {
		return errors.New("filetype is not supported")
	}

	return nil
}
