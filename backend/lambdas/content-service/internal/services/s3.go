package services

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/google/uuid"
)

type S3Service interface {
	UploadFile(file multipart.File, header *multipart.FileHeader, userID string) (string, error)
	GeneratePresignedURL(key string, expiration time.Duration) (string, error)
	DeleteFile(key string) error
}

type s3Service struct {
	client *s3.S3
	bucket string
}

func NewS3Service() S3Service {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))

	return &s3Service{
		client: s3.New(sess),
		bucket: os.Getenv("S3_BUCKET_NAME"),
	}
}

func (s *s3Service) UploadFile(file multipart.File, header *multipart.FileHeader, userID string) (string, error) {
	defer file.Close()

	// Validate file type
	if !isValidFileType(header.Filename) {
		return "", fmt.Errorf("invalid file type. Only PDF files are allowed")
	}

	// Generate unique filename
	ext := filepath.Ext(header.Filename)
	filename := fmt.Sprintf("%s/%s%s", userID, uuid.New().String(), ext)

	// Read file content
	buffer := make([]byte, header.Size)
	_, err := file.Read(buffer)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	// Upload to S3
	_, err = s.client.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(s.bucket),
		Key:           aws.String(filename),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(header.Size),
		ContentType:   aws.String("application/pdf"),
		Metadata: map[string]*string{
			"original-filename": aws.String(header.Filename),
			"user-id":           aws.String(userID),
		},
	})

	if err != nil {
		return "", fmt.Errorf("failed to upload file to S3: %v", err)
	}

	return filename, nil
}

func (s *s3Service) GeneratePresignedURL(key string, expiration time.Duration) (string, error) {
	req, _ := s.client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	url, err := req.Presign(expiration)
	if err != nil {
		return "", fmt.Errorf("failed to generate presigned URL: %v", err)
	}

	return url, nil
}

func (s *s3Service) DeleteFile(key string) error {
	_, err := s.client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		return fmt.Errorf("failed to delete file from S3: %v", err)
	}

	return nil
}

func isValidFileType(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".pdf"
}
