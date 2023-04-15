package services

import (
	"errors"
	"my-gram/models"
	"my-gram/repositories"
)

type CommentService interface {
	CreateComment(comment *models.Comment, userId uint) error
	UpdateComment(comment *models.Comment, userId uint, id uint) error
	DeleteComment(comment *models.Comment, userId uint, id uint) error
	GetCommentById(id uint, userId uint) (*models.Comment, error)
	GetAllComment() (*[]models.Comment, error)
}

type commentService struct {
	repo      repositories.CommentRepository
	photoRepo repositories.PhotoRepository
}

func NewCommentService(repo repositories.CommentRepository, photoRepo repositories.PhotoRepository) CommentService {
	return &commentService{repo: repo, photoRepo: photoRepo}
}

func (s *commentService) CreateComment(comment *models.Comment, userId uint) error {
	_, err := s.photoRepo.GetPhotoById(comment.PhotoId, userId)

	if err != nil {
		return errors.New("Photo with that id not found")
	}

	return s.repo.CreateComment(comment)
}

func (s *commentService) UpdateComment(comment *models.Comment, userId uint, id uint) error {
	comment.UserID = userId
	comment.ID = id
	return s.repo.UpdateComment(comment)
}

func (s *commentService) DeleteComment(comment *models.Comment, userId uint, id uint) error {
	comment.UserID = userId
	comment.ID = id
	return s.repo.DeleteComment(comment)
}

func (s *commentService) GetCommentById(id uint, userId uint) (*models.Comment, error) {
	return s.repo.GetCommentById(id, userId)
}

func (s *commentService) GetAllComment() (*[]models.Comment, error) {
	return s.repo.GetAllComment()
}
