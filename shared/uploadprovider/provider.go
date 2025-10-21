package uploadprovider

import (
	"context"
	"github.com/ngleanhvu/go-booking/shared/core"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, data []byte, dst string) (*core.Image, error)
}
