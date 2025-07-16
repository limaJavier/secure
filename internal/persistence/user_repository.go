package persistence

import (
	"fmt"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user User) error
	Retrieve(username string) (User, error)
}

func NewUserRepository() (UserRepository, error) {
	db, err := getDb(false)
	if err != nil {
		return nil, fmt.Errorf("cannot initialize user-repository: %v", err)
	}

	return &userRepository{
		db: db,
	}, nil
}

type userRepository struct {
	db *gorm.DB
}

func (repository *userRepository) Create(user User) error {
	result := repository.db.Create(&user)
	if result.Error != nil {
		return fmt.Errorf("cannot create user: %v", result.Error)
	}
	return nil
}

func (repository *userRepository) Retrieve(username string) (User, error) {
	var user User
	result := repository.db.First(&user, "username = ?", username)
	if result.Error != nil {
		return User{}, fmt.Errorf("cannot retrieve user: %v", result.Error)
	}
	return user, nil
}
