package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type PhotoRepository interface {
	CreatePhoto(photo *models.Photo) error
	UpdatePhoto(photo *models.Photo) error
	DeletePhoto(photo *models.Photo) error
	GetPhotoById(id uint, userId uint) (*models.Photo, error)
	GetAllPhoto() (*[]models.Photo, error)
}

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) PhotoRepository {
	return &photoRepository{db: db}
}

func (r *photoRepository) CreatePhoto(photo *models.Photo) error {
	return r.db.Create(photo).Error
}

func (r *photoRepository) UpdatePhoto(photo *models.Photo) error {
	return r.db.Model(photo).Updates(models.Photo{Title: photo.Title, Caption: photo.Caption, PhotoUrl: photo.PhotoUrl}).Error
}

func (r *photoRepository) DeletePhoto(photo *models.Photo) error {
	return r.db.Delete(photo, "id = ?", photo.ID).Error
}

func (r *photoRepository) GetPhotoById(id uint, userId uint) (*models.Photo, error) {
	photo := &models.Photo{}
	err := r.db.First(photo, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return photo, nil
}

func (r *photoRepository) GetAllPhoto() (*[]models.Photo, error) {
	photos := &[]models.Photo{}

	err := r.db.Find(&photos).Error
	if err != nil {
		return nil, err
	}
	return photos, nil
}
