package uploadprovider

import (
	"bytes"
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
		bucketName:   bucketName,
		accountId:    accountId,
		accessKey:    accessKey,
		secretKey:    secretKey,
		domain: domain,
	}

	sess, err := session.NewSession(&aws.Config{
		Region:           aws.String("auto"), // R2 không cần region thật
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
func (provider *r2Provider) SaveFileUploaded(ctx context.Context, file io.Reader, dst string) (*core.Image, error) {
	buf := make([]byte, 512)
    n, _ := file.Read(buf)

	file = io.MultiReader(bytes.NewReader(buf[:n]), file)

	uploader := s3manager.NewUploader(provider.session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(provider.bucketName),
		Key:    aws.String(dst),
		ACL:    aws.String("public-read"),
		Body:   file, 
	})
	
	//req, _ := s3.New(provider.session).PutObjectRequest(&s3.PutObjectInput{
	//	Bucket: aws.String(provider.bucketName),
	//	Key:    aws.String(dst),
	//	ACL:    aws.String("private"),
	//})
	//
	//req.Presign(time.Second * 5)

	if err != nil {
		return nil, err
	}

	img := &core.Image{}

	fmt.Printf("Successfully uploaded %q to %q\n", dst, provider.bucketName)
   
	return img, nil
}
