package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joaocansi/simple-api/internal/config"
)

type TokenService struct {
	secretKey string
	expiredIn uint
}

func NewTokenService(config *config.Config) *TokenService {
	return &TokenService{config.Token.SecretKey, uint(config.Token.ExpiresIn)}
}

func (t *TokenService) GenerateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userID,
		"exp": jwt.NewNumericDate(time.Now().Add(time.Duration(t.expiredIn) * time.Second)),
	})
	signedToken, err := token.SignedString([]byte(t.secretKey))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (t *TokenService) Validate(token string) (jwt.MapClaims, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(jwt.MapClaims); ok {
		return claims, nil
	} else {
		return nil, err
	}
}
