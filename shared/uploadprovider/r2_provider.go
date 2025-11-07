package uploadprovider

import (
	"context"
	"fmt"
	"io"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type r2Provider struct {
	bucketName string
	accountId  string
	accessKey  string
	secretKey  string
	domain     string
	session    *session.Session
}

func NewR2Provider(bucketName, accountId, accessKey, secretKey, domain string) *r2Provider {
	endpoint := fmt.Sprintf("https://%s.r2.cloudflarestorage.com", accountId)

	provider := &r2Provider{
		bucketName: bucketName,
		accountId:  accountId,
		accessKey:  accessKey,
		secretKey:  secretKey,
		domain:     domain,
	}

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("auto"),
		Endpoint:         aws.String(endpoint),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
	})

	if err != nil {
		panic(err)
	}

	provider.session = sess
	return provider
}

// SaveFileUploaded receives data and stores it into aws s3
func (provider *r2Provider) SaveFileUploaded(ctx context.Context, file io.Reader, dst, contentType string) (*core.Image, error) {

	// fileBytes, err := io.ReadAll(file)
	// if err != nil {
	// 	return nil, fmt.Errorf("failed to read file: %w", err)
	// }
	
	// fileReader := bytes.NewReader(fileBytes)

	uploader := s3manager.NewUploader(provider.session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket:      aws.String(provider.bucketName),
		Key:         aws.String(dst),
		ACL:         aws.String("public-read"),
		Body:        file,
		ContentType: aws.String(contentType),
	})

	if err != nil {
		return nil, err
	}

	img := &core.Image{}
	img.Fulfill(provider.domain, dst)
	fmt.Printf("Successfully uploaded %q to %q\n", dst, "R2")

	return img, nil
}
