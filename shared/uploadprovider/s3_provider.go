package uploadprovider

import (
	"bytes"
	"context"
	"fmt"
	"image"
	"io"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/ngleanhvu/go-booking/shared/core"

	"log"
)

type s3Provider struct {
	bucketName string
	region     string
	apiKey     string
	secret     string
	domain     string
	session    *session.Session
}

func NewS3Provider(bucketName string, region string, apiKey string, secret string, domain string) *s3Provider {
	provider := &s3Provider{
		bucketName: bucketName,
		region:     region,
		apiKey:     apiKey,
		secret:     secret,
		domain:     domain,
	}
	s3Session, err := session.NewSession(&aws.Config{
		Region: aws.String(provider.region),
		Credentials: credentials.NewStaticCredentials(
			provider.apiKey, // Access key ID
			provider.secret, // Secret access key
			""),             // Token
	})
	if err != nil {
		log.Fatal(err)
	}
	provider.session = s3Session
	return provider
}

// SaveFileUploaded receives data and stores it into aws s3
func (provider *s3Provider) SaveFileUploaded(ctx context.Context, file io.Reader, dst, contentType string) (*core.Image, error) {
	buf := make([]byte, 512)
	n, _ := file.Read(buf)

	file = io.MultiReader(bytes.NewReader(buf[:n]), file)

	uploader := s3manager.NewUploader(provider.session)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(provider.bucketName),
		Key:    aws.String(dst),
		ACL:    aws.String("private"),
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

func HandleImage(ctx context.Context,
	folder string,
	fileHeader *multipart.FileHeader,
	up UploadProvider) (*core.Image, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return nil, core.ErrUploadFile
	}
	defer func(file multipart.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(file); err != nil {
		return nil, core.ErrUploadFile
	}

	config, format, err := image.DecodeConfig(bytes.NewReader(buf.Bytes()))
	if err != nil {
		return nil, core.ErrUploadFile
	}

	filename := fmt.Sprintf("%d.%s", time.Now().UnixNano(), format)

	fmt.Printf("Image format: %s, width: %d, height: %d\n", format, config.Width, config.Height)

	contentType := fileHeader.Header.Get("Content-Type")
	dst := filepath.Join(folder, filename)

	img, err := up.SaveFileUploaded(ctx, file, dst, contentType)

	if err != nil {
		return nil, core.ErrUploadFile
	}

	img.Width = config.Width
	img.Height = config.Height

	return img, nil
}
