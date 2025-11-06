package facilitybiz

import (
	"context"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"path/filepath"
	"time"

	"github.com/ngleanhvu/go-booking/services/property/module/facility/model"
	"github.com/ngleanhvu/go-booking/shared/core"
)

func (b *facilityBusiness) CreateFacilityBiz(ctx context.Context, data *model.FacilityCreateDto, file io.Reader, folder string) error {
	config, format, err := image.DecodeConfig(file)
	
	if err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrInvalidIcon.Error()).
			WithDebug(err.Error())
	}

	filename := fmt.Sprintf("%d.%s", time.Now().Nanosecond(), format)
	dst := filepath.Join(folder, filename)
	
	img, err := b.uploader.SaveFileUploaded(ctx, file, dst); 
	
	if err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotUploadIcon.Error()).
			WithDebug(err.Error())
	}

	img.Width = config.Width
	img.Height = config.Height

	data.Icon = img

	if err := b.repo.Create(ctx, data); err != nil {
		return core.ErrInternalServerError.
			WithError(model.ErrCannotCreateFacility.Error()).
			WithDebug(err.Error())
	}

	return nil
}