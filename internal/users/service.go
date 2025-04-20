package users

import (
	"github.com/joaocansi/simple-api/internal/helpers/errors"
	"github.com/joaocansi/simple-api/internal/helpers/hash"
	"github.com/joaocansi/simple-api/internal/token"
	model "github.com/joaocansi/simple-api/storage/model"
	repository "github.com/joaocansi/simple-api/storage/repository"
	"gorm.io/gorm"
)

type UserService struct {
	userRepository *repository.UserRepository
	tokenService *token.TokenService
}

type CreateUser struct {
	Name      string
	Email     string
	Password  string
	AvatarUrl string
}

type SignIn struct {
	Email    string `json:"email"` 
	Password string `json:"password"`
}

func NewUserService(db *gorm.DB, tokenService *token.TokenService) *UserService {
	userRepository := repository.NewUserRepository(db)
	return &UserService{userRepository, tokenService}
}

func (s *UserService) createUser(data CreateUser) (*model.User, *errors.ServiceError) {
	_, err := s.userRepository.GetByEmail(data.Email)
	if err == nil {
		return nil, errors.UserAlreadyExists()
	}

	hashedPassword, err := hash.Hash(data.Password)
	if err != nil {
		return nil, errors.InternalError()
	}

	user, err := s.userRepository.Store(repository.CreateUserData{
		Name:      data.Name,
		Email:     data.Email,
		Password:  hashedPassword,
		AvatarUrl: data.AvatarUrl,
	})

	if err != nil {
		return nil, errors.InternalError()
	}

	return user, nil
}

type SignInResult struct {
	AccessToken string  `json:"accessToken"`
}

func (s *UserService) signIn(data SignIn) (*SignInResult, *errors.ServiceError) {
	user, _ := s.userRepository.GetByEmail(data.Email)
	if user == nil {
		return nil, errors.UserNotFound()
	}

	if err := hash.Verify(data.Password, user.Password); err != nil {
		return nil, errors.WrongUserCredentials()
	}

	token, err := s.tokenService.GenerateToken(user.ID);
	if err != nil {
		return nil, errors.WrongUserCredentials()
	}

	return &SignInResult{
		AccessToken: token,
	}, nil
}
