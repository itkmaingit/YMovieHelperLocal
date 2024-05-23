package utils

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func UploadToS3(ctx context.Context, fileName string, data io.Reader) (string, error) {

	var fileUrl string

	bucket := os.Getenv("AWS_S3_BUCKET_NAME")
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fileUrl, fmt.Errorf("UploadToS3: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   data,
	})
	if err != nil {
		return fileUrl, fmt.Errorf("UploadToS3.PutObject: %w", err)
	}

	fileUrl = fmt.Sprintf("%s/%s", os.Getenv("AWS_CLOUDFRONT_DOMAIN"), fileName)

	return fileUrl, err
}

func DownloadFromCroudFront(fileUrl string) (data []byte, ext string, err error) {
	resp, err := http.Get(fileUrl)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	data, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, "", err
	}

	ext = strings.ToLower(filepath.Ext(fileUrl))

	return data, ext, nil
}

func DeleteOnS3(ctx context.Context, fileUrl string) error {
	bucket := os.Getenv("AWS_S3_BUCKET_NAME")
	key := strings.Replace(fileUrl, os.Getenv("AWS_CLOUDFRONT_DOMAIN")+`/`, "", -1)
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return fmt.Errorf("DeleteOnS3: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return fmt.Errorf("DeleteOnS3: %w", err)
	}

	return nil
}

func RenameOnS3(ctx context.Context, oldFileName string, newFileName string) (string, error) {
	bucket := os.Getenv("AWS_S3_BUCKET_NAME")
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("RenameOnS3: %v", err)
	}

	client := s3.NewFromConfig(cfg)
	copySource := url.PathEscape(fmt.Sprintf("%s/%s", bucket, oldFileName))

	_, err = client.CopyObject(ctx, &s3.CopyObjectInput{
		Bucket:     aws.String(bucket),
		CopySource: aws.String(copySource),
		Key:        aws.String(newFileName),
	})
	if err != nil {
		return "", fmt.Errorf("RenameOnS3: %v", err)
	}

	_, err = client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(oldFileName),
	})

	if err != nil {
		return "", fmt.Errorf("RenameOnS3: %w", err)
	}

	fileUrl := fmt.Sprintf("%s/%s", os.Getenv("AWS_CLOUDFRONT_DOMAIN"), newFileName)

	return fileUrl, nil
}
