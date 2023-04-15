package repositories

import (
	"my-gram/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (ur *userRepository) CreateUser(user *models.User) error {
	return ur.db.Debug().Create(user).Error
}

func (ur *userRepository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := ur.db.Debug().Where("email = ?", email).Take(user).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}
