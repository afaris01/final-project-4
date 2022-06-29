package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

const (
	AdminRole  = "admin"
	MemberRole = "customer"
)

type User struct {
	ID       int    `gorm:"primaryKey" json:"id"`
	FullName string `gorm:"not null;uniqueIndex" json:"full_name" form:"full_name" valid:"required~Title is required"`
	Email    string `gorm:"not null;uniqueIndex" json:"email" form:"email" valid:"required~Email is required,email~Email is invalid"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,minstringlength(6)~Password has to have minimum length 6 characters"`
	Role     string `gorm:"not null;default:'member'" json:"role" form:"role" valid:"required~Role is requiered"`
	Balance  int    `json:"balance"`
	TimeModel
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
