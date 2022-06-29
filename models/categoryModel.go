package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Category struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Type      string    `gorm:"not null;uniqueIndex" json:"type" form:"type" valid:"required~Type is Required"`
	TimeModel
}

func (Category) TableName() string {
	return "category"
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
