package helper

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/ayyaa/todo-services/config"
	"github.com/ayyaa/todo-services/models"
)

func NewSession(awsCred config.AWSCredential) (*s3.S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(awsCred.Region),
		Credentials: credentials.NewStaticCredentials(
			awsCred.KeyID,
			awsCred.AccesKey,
			"",
		),
	})

	fmt.Println(err)

	if err != nil {
		return nil, err
	}

	svc := s3.New(sess)

	return svc, nil
}

func Upload(ctx context.Context, list *models.ListRequest, awsCred config.AWSCredential) ([]*models.Attachment, error) {
	attachments := []*models.Attachment{}

	for _, file := range list.File {
		// Source
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		attachment := models.Attachment{
			Filename: filename,
			Filepath: "upload/" + filename,
		}

		// Create an S3 object to be uploaded
		upParams := &s3.PutObjectInput{
			Bucket: aws.String(awsCred.BucketName),
			Key:    aws.String(attachment.Filepath),
			Body:   src,
		}

		// Upload the file to S3
		if _, err := awsCred.AWSSession.PutObject(upParams); err != nil {
			awsCred.AwsFlag = false
			if !awsCred.AwsFlag {
				return UploadLocal(ctx, list)
			}
		}

		attachments = append(attachments, &attachment)

	}

	return attachments, nil
}

func UploadLocal(ctx context.Context, list *models.ListRequest) ([]*models.Attachment, error) {

	attachments := []*models.Attachment{}

	for _, file := range list.File {

		// Source
		src, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer src.Close()

		filename := fmt.Sprintf("%d-%s", time.Now().Unix(), file.Filename)
		attachment := models.Attachment{
			Filename: filename,
			Filepath: filename,
		}

		// Create file
		dst, err := os.Create(file.Filename)
		defer dst.Close()
		if err != nil {
			return nil, err
		}

		if _, err := io.Copy(dst, src); err != nil {
			return nil, err
		}

		attachments = append(attachments, &attachment)

	}

	return attachments, nil
}
