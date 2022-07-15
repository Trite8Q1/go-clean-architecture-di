package repo

import (
	"github.com/trite8q1/go-clean-architecture-di/pkg/entity"
	"gorm.io/gorm"
)

// UserRepository : represent the user's repository contract
type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Delete(user *entity.User) error
	Migrate() error
}

type userRepository struct {
	DB *gorm.DB
}

// NewUserRepository : get injected database
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}
func (u *userRepository) Save(user *entity.User) (*entity.User, error) {
	return user, u.DB.Create(user).Error
}
func (u *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := u.DB.Find(&users).Error
	return users, err
}
func (u *userRepository) Delete(user *entity.User) error {
	return u.DB.Delete(&user).Error
}
func (u *userRepository) Migrate() error {
	return u.DB.AutoMigrate(&entity.User{})
}
