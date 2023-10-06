package pkg

import (
	"context"
	"ecommerce-backend/src/configs"
	"fmt"
	"go/types"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Service interface {
	GetS3Client(cfg *configs.Env) (*s3.Client, error)
	Download(bucket string, path *string) (*string, error)
	ListObjects(bucket string) ([]types.Object, error)
	UploadFile(bucket string, file *multipart.FileHeader) (*string, error)
}

type s3Service struct {
}

func NewS3Service() S3Service {
	return s3Service{}
}

func (s s3Service) GetS3Client(cfg *configs.Env) (*s3.Client, error) {
	r2Resolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL: fmt.Sprintf("https://%s.r2.cloudflarestorage.com", cfg.S3.AccountID),
		}, nil
	})

	cfgS3, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(r2Resolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(cfg.S3.AccessKeyID, cfg.S3.AccessKeySecret, "")),
	)
	if err != nil {
		return nil, err
	}
	return s3.NewFromConfig(cfgS3), nil
}

func (s s3Service) Download(bucket string, path *string) (*string, error) {
	return nil, nil
}

func (s s3Service) ListObjects(bucket string) ([]types.Object, error) {
	return nil, nil
}

func (s s3Service) UploadFile(bucket string, file *multipart.FileHeader) (*string, error) {
	return nil, nil
}
