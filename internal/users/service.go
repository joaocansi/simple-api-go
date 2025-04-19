package users

import (
	"fmt"

	"github.com/joaocansi/simple-api/internal/hash"
	storage "github.com/joaocansi/simple-api/internal/storage/model"
	"github.com/joaocansi/simple-api/internal/storage/repository"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository *repository.UserRepository
}

type CreateUser struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	AvatarUrl string `json:"avatarUrl"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUserService(db *gorm.DB) *UserService {
	userRepository := repository.NewUserRepository(db)
	return &UserService{userRepository}
}

func (s *UserService) createUser(data CreateUser) (*storage.User, error) {
	_, err := s.userRepository.GetByEmail(data.Email)
	if err == nil {
		return nil, fmt.Errorf("email %s já cadastrado", data.Email)
	}

	hashedPassword, err := hash.Hash(data.Password)
	if err != nil {
		return nil, fmt.Errorf("não foi possível criar usuário")
	}

	user, err := s.userRepository.Store(repository.CreateUserData{
		Name:      data.Name,
		Email:     data.Email,
		Password:  hashedPassword,
		AvatarUrl: data.AvatarUrl,
	})

	if err != nil {
		return nil, fmt.Errorf("não foi possível criar usuário")
	}

	return user, nil
}

type SignInResult struct {
	AccessToken string  `json:"accessToken"`
	ExpiresIn   float32 `json:"expiresIn"`
}

func (s *UserService) signIn(data SignIn) (*SignInResult, error) {
	user, _ := s.userRepository.GetByEmail(data.Email)
	if user == nil {
		return nil, fmt.Errorf("email ou senha não está correto")
	}

	if err := hash.Verify(data.Password, user.Password); err != nil {
		return nil, fmt.Errorf("email ou senha não está correto")
	}

	return &SignInResult{
		AccessToken: "Teste",
		ExpiresIn:   1000000,
	}, nil
}
