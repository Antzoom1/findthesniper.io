package s3

import (
	"context"

	"github.com/RagOfJoes/findthesniper.io/internal/config"
	"github.com/RagOfJoes/findthesniper.io/repositories"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsS3 "github.com/aws/aws-sdk-go-v2/service/s3"
)

var _ repositories.Uploader = (*s3)(nil)

type s3 struct {
	client        *awsS3.Client
	presignClient *awsS3.PresignClient
}

func New(cfg config.Configuration) (*s3, error) {
	client := awsS3.NewFromConfig(aws.Config{
		Region: "",
	})
	presignClient := awsS3.NewPresignClient(client)

	return &s3{
		client:        client,
		presignClient: presignClient,
	}, nil
}

func (s *s3) Verify(ctx context.Context, url string) error {
	panic("unimplemented")
}

func (s *s3) PresignUpload(ctx context.Context, fileKey string) (string, error) {
	panic("unimplemented")
}

func (s *s3) Delete(ctx context.Context, url string) error {
	panic("unimplemented")
}
