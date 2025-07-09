package model

import (
	"context"
	"fmt"
	"log"
	"mime/multipart"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const PAGINATED_COUNT = 15

func NewS3Config() (*s3.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("US"))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
	}), nil
}

func UploadTrackToS3(trackName string, data multipart.File) (*s3.PutObjectOutput, error) {
	svc, err := NewS3Config()
	if err != nil {
		return nil, err
	}

	o, err := svc.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String("inbox"),
		Key:    aws.String("track/" + trackName),
		Body:   data,
	})
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return o, nil
}
