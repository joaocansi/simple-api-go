package storage

import (
	model "github.com/joaocansi/simple-api/storage/model"
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

func (r *UserRepository) Store(data CreateUserData) (*model.User, error) {
	user := model.User{
		Name:      data.Name,
		Email:     data.Email,
		Password:  data.Password,
		AvatarUrl: data.AvatarUrl,
	}

	err := r.db.Create(&user).Error
	return &user, err
}

func (r *UserRepository) Update(data *model.User) (*model.User, error) {
	err := r.db.Save(data).Error
	return data, err
}

func (r *UserRepository) Get(id string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) GetByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, "email = ?", email).Error
	return &user, err
}
