package services

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/HeadGardener/game-tracker/internal/app/models"
	"github.com/HeadGardener/game-tracker/internal/app/repositories"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	salt      = "qetuoadgjlzcbmwryipsfhkxvn"
	tokenTTL  = 2 * time.Hour
	secretKey = "qazwsxedcrfvtgbyhnujm"
)

type AuthService struct {
	repos *repositories.Repository
}

func NewAuthService(repos *repositories.Repository) *AuthService {
	return &AuthService{repos: repos}
}

type tokenClaims struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

func (s *AuthService) Create(userInput models.RegUserInput) (int, error) {
	user := models.User{
		Name:         userInput.Name,
		Username:     userInput.Username,
		PasswordHash: getPasswordHash(userInput.Password),
	}

	return s.repos.Authorization.Create(user)
}

func (s *AuthService) GenerateToken(userInput models.LogUserInput) (string, error) {
	user := models.User{
		Username:     userInput.Username,
		PasswordHash: getPasswordHash(userInput.Password),
	}

	err := s.repos.GetUser(&user)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(secretKey))
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}

func getPasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
