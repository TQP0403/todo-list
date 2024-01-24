package file

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"errors"
	"mime/multipart"
	"os"
	"slices"
	"strings"
)

var allowTypes = []string{"image/jpeg", "image/png", "image/webp"}

var fileLimit = helper.GetDefaultNumber[int64](
	int64(helper.ParseInt(os.Getenv("UPLOAD_FILE_LIMIT"))),
	5*1024*1024, // 5mb
)

type ImageFile struct {
	Name string
	Body multipart.File
}

func ReadFile(fileHeader *multipart.FileHeader) (*ImageFile, error) {
	f := &ImageFile{}
	f.Name = strings.ReplaceAll(fileHeader.Filename, " ", "-")

	if err := validateUploadFile(fileHeader); err != nil {
		return nil, common.NewBadRequestError(err)
	}

	if file, err := fileHeader.Open(); err != nil {
		return nil, common.NewBadRequestError(err)
	} else {
		f.Body = file
		return f, nil
	}
}

func validateUploadFile(file *multipart.FileHeader) error {
	if file.Size > fileLimit {
		return errors.New("file too large")
	}

	if !slices.Contains(allowTypes, file.Header.Get("Content-Type")) {
		return errors.New("filetype is not supported")
	}

	return nil
}
