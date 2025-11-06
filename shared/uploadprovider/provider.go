package uploadprovider

import (
	"context"
	"io"

	"github.com/ngleanhvu/go-booking/shared/core"
)

type UploadProvider interface {
	SaveFileUploaded(ctx context.Context, file io.Reader, dst string) (*core.Image, error)
}
