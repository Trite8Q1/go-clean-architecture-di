package service

import (
	"errors"
	"sync"
	"time"

	"github.com/trite8q1/go-clean-architecture-di/internal/user/repo"
	"github.com/trite8q1/go-clean-architecture-di/pkg/entity"
)

type UserService interface {
	Validate(user *entity.User) error
	ValidateAge(user *entity.User) bool
	Create(user *entity.User) (*entity.User, error)
	FindAll() ([]entity.User, error)
}

var once sync.Once

type userService struct {
	userRepository repo.UserRepository
}

var instance *userService

//NewUserService: construction function, injected by user repository
func NewUserService(r repo.UserRepository) UserService {
	once.Do(func() {
		instance = &userService{
			userRepository: r,
		}
	})
	return instance
}
func (*userService) Validate(user *entity.User) error {
	if user == nil {
		err := errors.New("the user is empty")
		return err
	}
	if user.Name == "" {
		err := errors.New("the name of user is empty")
		return err
	}
	if user.Email == "" {
		err := errors.New("the email of user is empty")
		return err
	}
	if user.DOB == "" {
		err := errors.New("tehe DOB of user is empty")
		return err
	}
	return nil
}
func (*userService) ValidateAge(user *entity.User) bool {
	ageLimit := 13
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	dob, err := time.Parse("2006-01-02", user.DOB)
	if err != nil {
		return false
	}
	diff := now.Sub(dob)
	diffInYears := int(diff.Hours() / (24 * 7 * 4 * 12))
	if diffInYears < ageLimit {
		return false
	} else {
		return true
	}
}
func (u *userService) Create(user *entity.User) (*entity.User, error) {
	return u.userRepository.Save(user)
}
func (u *userService) FindAll() ([]entity.User, error) {
	return u.userRepository.FindAll()
}
