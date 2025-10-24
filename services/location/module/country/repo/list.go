package repo

import (
	"context"
	"fmt"
	"strings"

	"github.com/ngleanhvu/go-booking/services/location/module/country/model"
	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

func (s *postgresRepo) List(ctx context.Context,
	filter *model.Filter,
	paging *core.Paging,
	moreKeys ...string,
) ([]model.Country, error) {
	var result []model.Country

	db := s.db.Table(model.Country{}.TableName())

	if f := filter; f != nil {
		keyword := strings.TrimSpace(f.Keyword)
		if keyword != "" {
			db = db.Where("lower(name) LIKE ?", keyword)
		}
	}

	var total int64

	if err := db.Count(&total).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	paging.Total = int(total)

	if v := paging.FakeCursor; v != "" {
		uid, err := core.FromBase58(v)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		db = db.Order("id desc").Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1)*paging.Limit
		db = db.Offset(offset).Order(fmt.Sprintf("%s %s", paging.))
	}

}
