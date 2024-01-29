package file

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"context"
	"errors"
	"log"
	"mime/multipart"
	"net/url"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type IFileService interface {
	UploadFile(header *multipart.FileHeader) (string, error)
}

type FileService struct {
	enable       bool
	validator    IFileValidator
	s3Cloudfront string
	s3Bucket     string
	s3Client     *s3.Client
}

func NewService() *FileService {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	s3Cloudfront := os.Getenv("AWS_S3_CLOUD_FRONT")
	s3Bucket := os.Getenv("AWS_S3_BUCKET")
	enable := err == nil && len(s3Cloudfront) > 0 && len(s3Bucket) > 0
	if err != nil {
		log.Println(err)
	}
	if strings.LastIndex(s3Cloudfront, "https://") == -1 {
		s3Cloudfront = "https://" + s3Cloudfront
	}

	return &FileService{
		enable:       enable,
		validator:    NewValidator(),
		s3Cloudfront: s3Cloudfront,
		s3Bucket:     s3Bucket,
		s3Client:     s3.NewFromConfig(cfg),
	}
}

func (service *FileService) UploadFile(fileHeader *multipart.FileHeader) (string, error) {
	// service enable check
	if !service.enable {
		return "", common.NewServiceUnavailableError(errors.New("upload file service disable"))
	}

	// validate upload file
	if err := service.validator.Validate(fileHeader); err != nil {
		return "", common.NewBadRequestError(err)
	}

	// open file
	file, err := fileHeader.Open()
	if err != nil {
		return "", common.NewBadRequestError(err)
	}
	defer file.Close()

	// store on AWS S3
	key := helper.RandomAplphaNumeric(16) + "-" + strings.ReplaceAll(fileHeader.Filename, " ", "-")

	if _, err = service.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket:      aws.String(service.s3Bucket),
		Key:         aws.String(key),
		ContentType: aws.String(fileHeader.Header.Get("Content-Type")),
		Body:        file,
	}); err != nil {
		return "", common.NewInternalServerError(err)
	}

	return url.JoinPath(service.s3Cloudfront, key)
}
