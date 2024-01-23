package file

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

var fileLimit = helper.GetDefaultNumber[int64](
	int64(helper.ParseInt(os.Getenv("UPLOAD_FILE_LIMIT"))),
	1024*1024, // 1mb
)

type FileService struct{}

type IFileService interface {
	UploadFile(fileHeader *multipart.FileHeader) (string, error)
}

func NewService() *FileService {
	return &FileService{}
}

func (service *FileService) UploadFile(header *multipart.FileHeader) (string, error) {
	fileName, _, err := readFile(header)
	if err != nil {
		return "", err
	}

	key := helper.RandomAplphaNumeric(16) + "-" + fileName

	// Read the contents of the file into a buffer

	// // This uploads the contents of the buffer to S3
	// _, err := svc.PutObject(&s3.PutObjectInput{
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String(key),
	// 	Body:   bytes.NewReader(buffer.Bytes()),
	// })
	// if err != nil {
	// 	return "", common.NewInternalServerError(err)
	// }

	fileUrl := key

	return fileUrl, nil
}

func readFile(fileHeader *multipart.FileHeader) (string, bytes.Buffer, error) {
	fileName := strings.ReplaceAll(fileHeader.Filename, " ", "-")
	var buffer bytes.Buffer

	if err := validateUploadFile(fileHeader); err != nil {
		return "", buffer, common.NewBadRequestError(err)
	}

	file, err := fileHeader.Open()
	if err != nil {
		return "", buffer, common.NewBadRequestError(err)
	}
	// close file
	defer file.Close()

	if _, err := io.Copy(&buffer, file); err != nil {
		return "", buffer, common.NewInternalServerError(err)
	}

	return fileName, buffer, nil
}

func validateUploadFile(file *multipart.FileHeader) error {
	if file.Size > fileLimit {
		return errors.New("file too large")
	}

	if contentType := file.Header.Get("Content-Type"); contentType != "image/jpeg" && contentType != "image/png" {
		return errors.New("filetype is not supported")
	}

	return nil
}
