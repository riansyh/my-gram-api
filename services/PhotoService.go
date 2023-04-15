package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type PhotoService interface {
	CreatePhoto(photo *models.Photo) error
	UpdatePhoto(photo *models.Photo, userId uint, id uint) error
	DeletePhoto(photo *models.Photo, userId uint, id uint) error
	GetPhotoById(id uint, userId uint) (*models.Photo, error)
	GetAllPhoto() (*[]models.Photo, error)
}

type photoService struct {
	repo repositories.PhotoRepository
}

func NewPhotoService(repo repositories.PhotoRepository) PhotoService {
	return &photoService{repo: repo}
}

func (s *photoService) CreatePhoto(photo *models.Photo) error {
	return s.repo.CreatePhoto(photo)
}

func (s *photoService) UpdatePhoto(photo *models.Photo, userId uint, id uint) error {
	photo.UserID = userId
	photo.ID = id
	return s.repo.UpdatePhoto(photo)
}

func (s *photoService) DeletePhoto(photo *models.Photo, userId uint, id uint) error {
	photo.UserID = userId
	photo.ID = id
	return s.repo.DeletePhoto(photo)
}

func (s *photoService) GetPhotoById(id uint, userId uint) (*models.Photo, error) {
	return s.repo.GetPhotoById(id, userId)
}

func (s *photoService) GetAllPhoto() (*[]models.Photo, error) {
	return s.repo.GetAllPhoto()
}
