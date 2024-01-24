package file

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"context"
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
	s3Cloudfront string
	s3Bucket     string
	s3Client     *s3.Client
}

func NewService() *FileService {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Println(err)
	}

	s3Cloudfront := os.Getenv("S3_CLOUD_FRONT")
	s3Bucket := os.Getenv("S3_BUCKET")
	enable := err == nil && len(s3Cloudfront) > 0 && len(s3Bucket) > 0
	if strings.LastIndex(s3Cloudfront, "https://") == -1 {
		s3Cloudfront = "https://" + s3Cloudfront
	}

	return &FileService{
		enable:       enable,
		s3Cloudfront: s3Cloudfront,
		s3Bucket:     s3Bucket,
		s3Client:     s3.NewFromConfig(cfg),
	}
}

func (service *FileService) UploadFile(header *multipart.FileHeader) (string, error) {
	if !service.enable {
		return "", nil
	}

	f, err := ReadFile(header)
	if err != nil {
		return "", err
	}
	defer f.Body.Close()

	key := helper.RandomAplphaNumeric(16) + "-" + f.Name

	_, err = service.s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(""),
		Key:    aws.String(key),
		Body:   f.Body,
	})
	if err != nil {
		return "", common.NewInternalServerError(err)
	}

	return url.JoinPath(service.s3Cloudfront, key)
}
