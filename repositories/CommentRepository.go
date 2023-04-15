package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type CommentRepository interface {
	CreateComment(comment *models.Comment) error
	UpdateComment(comment *models.Comment) error
	DeleteComment(comment *models.Comment) error
	GetCommentById(id uint, userId uint) (*models.Comment, error)
	GetAllComment() (*[]models.Comment, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) CreateComment(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) UpdateComment(comment *models.Comment) error {
	return r.db.Model(comment).Updates(models.Comment{Message: comment.Message, UserID: comment.UserID}).Error
}

func (r *commentRepository) DeleteComment(comment *models.Comment) error {
	return r.db.Delete(comment, "id = ?", comment.ID).Error
}

func (r *commentRepository) GetCommentById(id uint, userId uint) (*models.Comment, error) {
	comment := &models.Comment{}
	err := r.db.First(comment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *commentRepository) GetAllComment() (*[]models.Comment, error) {
	comments := &[]models.Comment{}

	err := r.db.Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
