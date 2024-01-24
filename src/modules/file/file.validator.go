package file

import (
	"TQP0403/todo-list/src/helper"
	"errors"
	"mime/multipart"
	"os"
	"slices"
)

var allowTypes = []string{"image/jpeg", "image/png", "image/webp"}

var fileLimit = helper.GetDefaultNumber[int64](
	int64(helper.ParseInt(os.Getenv("UPLOAD_FILE_LIMIT"))),
	5*1024*1024, // 5mb
)

func ValidateUploadFile(file *multipart.FileHeader) error {
	if file.Size > fileLimit {
		return errors.New("file too large")
	}

	if !slices.Contains(allowTypes, file.Header.Get("Content-Type")) {
		return errors.New("filetype is not supported")
	}

	return nil
}
