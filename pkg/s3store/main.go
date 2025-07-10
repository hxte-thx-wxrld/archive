package s3store

import (
	"context"
	"embed"
	"fmt"
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

	anonPolicyCover, err := statics.ReadFile("assets/anon-policy-covers.json")
	if err != nil {
		log.Fatalln(err)
	}

	anonPolicyArtists, err := statics.ReadFile("assets/anon-policy-artists.json")
	if err != nil {
		log.Fatalln(err)
	}
	if !searchForBucket(result, "covers") {
		ctx := context.WithoutCancel(context.Background())

		data, err := statics.Open("assets/default-cover.png")
		if err != nil {
			log.Fatalln(err)
		}
		svc.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String("covers"),
		})
		_, err = svc.PutBucketPolicy(context.Background(), &s3.PutBucketPolicyInput{
			Bucket: aws.String("covers"),
			Policy: aws.String(string(anonPolicyCover)),
		})

		if err != nil {
			log.Fatalln(err)
		}

		UploadFileToS3(svc, "default.png", "covers", data)

	}

	if !searchForBucket(result, "artists") {
		ctx := context.WithoutCancel(context.Background())

		data, err := statics.Open("assets/default-artist.png")
		if err != nil {
			log.Fatalln(err)
		}
		svc.CreateBucket(ctx, &s3.CreateBucketInput{
			Bucket: aws.String("artists"),
		})

		_, err = svc.PutBucketPolicy(context.Background(), &s3.PutBucketPolicyInput{
			Bucket: aws.String("artists"),
			Policy: aws.String(string(anonPolicyArtists)),
		})
		if err != nil {
			log.Fatalln(err)
		}
		UploadFileToS3(svc, "default.png", "artists", data)
	}

	if !searchForBucket(result, "tracks") {
		svc.CreateBucket(context.TODO(), &s3.CreateBucketInput{
			Bucket: aws.String("tracks"),
		})

		/*_, err = svc.PutBucketPolicy(context.Background(), &s3.PutBucketPolicyInput{
			Bucket: aws.String("tracks"),
			Policy: aws.String(string(anonPolicy)),
		})
		if err != nil {
			log.Fatalln(err)
		}
		*/

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

func DeleteTrackFromInbox(filename string) error {
	svc, err := NewS3Config()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return DeleteFileFromS3(svc, "inbox", filename)
}
func DeleteFileFromS3(svc *s3.Client, filename string, bucket_name string) error {
	o, err := svc.DeleteObject(context.Background(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucket_name),
		Key:    aws.String(filename),
	})

	fmt.Println(o)
	return err
}
