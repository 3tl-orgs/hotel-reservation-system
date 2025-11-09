package propertydetailmodel

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/ngleanhvu/go-booking/shared/core"
	"github.com/pkg/errors"
)

type PropertyDetail struct {
	core.SQLModel
	PropertyId    int            `json:"property_id" gorm:"column:property_id"`
	CheckInTime   *string        `json:"check_in_time" gorm:"column:check_in_time"`
	CheckOutTime  *string        `json:"check_out_time" gorm:"column:check_out_time"`
	Description   *string        `json:"description" gorm:"column:description"`
	Hotline       *string        `json:"hotline" gorm:"column:hotline"`
	HostId        *int           `json:"host_id" gorm:"column:host_id"`
	WebsiteUrl    *string        `json:"website_url" gorm:"column:website_url"`
	Email         *string        `json:"email" gorm:"column:email"`
	Images        *core.Images   `json:"images" gorm:"column:images"`
	BestAmenities *BestAmenities `json:"best_amenities" gorm:"column:best_amenities"`
}

func (PropertyDetail) TableName() string {
	return "property_details"
}

type BestAmenities []int

func (img *BestAmenities) Scan(value interface{}) error { //Từ db -> scan -> trả ra value
	bytes, ok := value.([]byte)
	if !ok {
		return errors.WithStack(errors.New(fmt.Sprintf("Failed to unmarshal data from DB: %s", value)))
	}

	var i BestAmenities
	if err := json.Unmarshal(bytes, &i); err != nil {
		return errors.WithStack(err)
	}

	*img = i
	return nil
}

func (img *BestAmenities) Value() (driver.Value, error) {
	if img == nil {
		return nil, nil
	}
	return json.Marshal(img)
}
