package uploadprovider

import (
	"context"
	"go-booking/shared/core"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*core.Image, error)
}
