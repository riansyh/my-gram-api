package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

// SocialMedia represents the model for a social media
type SocialMedia struct {
	GormModel
	Name           string `gorm:"not null" json:"name" form:"name" valid:"required~Your name is required"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url" form:"social_media_url" valid:"required~Your social media url is required"`
	UserID         uint   `json:"user_id"`
	User           *User
}

func (s *SocialMedia) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
