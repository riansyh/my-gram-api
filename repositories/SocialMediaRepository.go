package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type SocialMediaRepository interface {
	CreateSocialMedia(socialMedia *models.SocialMedia) error
	UpdateSocialMedia(socialMedia *models.SocialMedia) error
	DeleteSocialMedia(socialMedia *models.SocialMedia) error
	GetSocialMediaById(id uint, userId uint) (*models.SocialMedia, error)
	GetAllSocialMedia() (*[]models.SocialMedia, error)
}

type socialMediaRepository struct {
	db *gorm.DB
}

func NewSocialMediaRepository(db *gorm.DB) SocialMediaRepository {
	return &socialMediaRepository{db: db}
}

func (r *socialMediaRepository) CreateSocialMedia(socialMedia *models.SocialMedia) error {
	return r.db.Create(socialMedia).Error
}

func (r *socialMediaRepository) UpdateSocialMedia(socialMedia *models.SocialMedia) error {
	return r.db.Model(socialMedia).Updates(models.SocialMedia{Name: socialMedia.Name, SocialMediaUrl: socialMedia.SocialMediaUrl}).Error
}

func (r *socialMediaRepository) DeleteSocialMedia(socialMedia *models.SocialMedia) error {
	return r.db.Delete(socialMedia, "id = ?", socialMedia.ID).Error
}

func (r *socialMediaRepository) GetSocialMediaById(id uint, userId uint) (*models.SocialMedia, error) {
	socialMedia := &models.SocialMedia{}
	err := r.db.First(socialMedia, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return socialMedia, nil
}

func (r *socialMediaRepository) GetAllSocialMedia() (*[]models.SocialMedia, error) {
	socialMedias := &[]models.SocialMedia{}

	err := r.db.Find(&socialMedias).Error
	if err != nil {
		return nil, err
	}
	return socialMedias, nil
}
