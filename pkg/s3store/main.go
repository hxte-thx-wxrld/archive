package s3store

import (
	"context"
	"embed"
	"io"
	"log"
	"mime/multipart"
	"sort"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func searchForBucket(target *s3.ListBucketsOutput, bucketName string) bool {
	_, found := sort.Find(len(target.Buckets), func(i int) int {
		return strings.Compare(bucketName, aws.ToString(target.Buckets[i].Name))
	})
	return found
}

func SetupS3Store(statics embed.FS) {
	svc, err := NewS3Config()
	if err != nil {
		log.Fatalln(err)
	}

	result, err := svc.ListBuckets(context.Background(), nil)
	if err != nil {
		log.Fatalln(err)
	}

	if !searchForBucket(result, "covers") {
		ctx := context.WithoutCancel(context.Background())

		data, err := statics.Open("assets/default.png")
		if err != nil {
			log.Fatalln(err)
		}
		svc.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String("covers"),
		})
		UploadFileToS3(svc, "default.png", "covers", data)

	}

	if !searchForBucket(result, "tracks") {
		svc.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String("tracks"),
		})

	}

	if !searchForBucket(result, "inbox") {
		svc.CreateBucket(context.Background(), &s3.CreateBucketInput{
			Bucket: aws.String("inbox"),
		})
	}
}

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

	return UploadFileToS3(svc, "track/"+trackName, "inbox", data)
}

func UploadFileToS3(svc *s3.Client, filename string, bucket_name string, data io.Reader) (*s3.PutObjectOutput, error) {
	return svc.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: aws.String(bucket_name),
		Key:    aws.String(filename),
		Body:   data,
	})
}
