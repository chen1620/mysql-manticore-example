package databases

import (
	"context"

	"gorm.io/gorm"

	"mysql-manticore-example/models"
)

// UserRepository database structure.
type UserRepository interface {
	CreateUser(ctx context.Context, u *models.User) error
	GetUserByID(ctx context.Context, id int) (*models.User, error)
}

// userRepo user repository implement struct.
type userRepo struct {
	db *gorm.DB
}

// NewUserRepository creates a new user repository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{
		db: db,
	}
}

// CreateUser ...
func (r *userRepo) CreateUser(ctx context.Context, u *models.User) error {
	//TODO implement me
	panic("implement me")
}

// GetUserByID ...
func (r *userRepo) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	var u models.User
	err := r.db.First(&u, id).Error
	return &u, err
}
