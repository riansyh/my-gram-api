package services

import (
	"my-gram/models"
	"my-gram/repositories"
)

type SocialMediaService interface {
	CreateSocialMedia(socialMedia *models.SocialMedia) error
	UpdateSocialMedia(socialMedia *models.SocialMedia, userId uint, id uint) error
	DeleteSocialMedia(socialMedia *models.SocialMedia, userId uint, id uint) error
	GetSocialMediaById(id uint, userId uint) (*models.SocialMedia, error)
	GetAllSocialMedia() (*[]models.SocialMedia, error)
}

type socialMediaService struct {
	repo repositories.SocialMediaRepository
}

func NewSocialMediaService(repo repositories.SocialMediaRepository) SocialMediaService {
	return &socialMediaService{repo: repo}
}

func (s *socialMediaService) CreateSocialMedia(socialMedia *models.SocialMedia) error {
	return s.repo.CreateSocialMedia(socialMedia)
}

func (s *socialMediaService) UpdateSocialMedia(socialMedia *models.SocialMedia, userId uint, id uint) error {
	socialMedia.UserID = userId
	socialMedia.ID = id
	return s.repo.UpdateSocialMedia(socialMedia)
}

func (s *socialMediaService) DeleteSocialMedia(socialMedia *models.SocialMedia, userId uint, id uint) error {
	socialMedia.UserID = userId
	socialMedia.ID = id
	return s.repo.DeleteSocialMedia(socialMedia)
}

func (s *socialMediaService) GetSocialMediaById(id uint, userId uint) (*models.SocialMedia, error) {
	return s.repo.GetSocialMediaById(id, userId)
}

func (s *socialMediaService) GetAllSocialMedia() (*[]models.SocialMedia, error) {
	return s.repo.GetAllSocialMedia()
}
