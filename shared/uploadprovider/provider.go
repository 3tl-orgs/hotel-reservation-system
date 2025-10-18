package uploadprovider

import (
	"context"
	"go-booking/shared/dto"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*dto.Image, error)
}
