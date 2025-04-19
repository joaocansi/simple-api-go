package repository

import (
	storage "github.com/joaocansi/simple-api/internal/storage/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type CreateUserData struct {
	Name      string
	Email     string
	Password  string
	AvatarUrl string
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Store(data CreateUserData) (*storage.User, error) {
	user := storage.User{
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		AvatarUrl: data.AvatarUrl,
	}

	err := r.db.Create(&user).Error
	return &user, err
}

func (r *UserRepository) Update(data *storage.User) (*storage.User, error) {
	err := r.db.Save(data).Error
	return data, err
}

func (r *UserRepository) Get(id string) (*storage.User, error) {
	var user storage.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByEmail(email string) (*storage.User, error) {
	var user storage.User
	err := r.db.First(&user, "email = ?", email).Error
	return &user, err
}
