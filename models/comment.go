package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// Comment represents the model for a comment
type Comment struct {
	GormModel
	Message string `gorm:"not null" json:"message" form:"message" valid:"required~The message is required"`
	PhotoId uint   `json:"photo_id"`
	UserID  uint   `json:"user_id"`
	User    *User
	Photo   *Photo
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
