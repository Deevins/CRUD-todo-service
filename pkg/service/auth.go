package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/pkg/repository"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

const (
	tokenTTL = 12 * time.Hour
)

var jwtKey = []byte(os.Getenv("SIGNING_KEY"))
var passwordSalt = []byte(os.Getenv("SALT"))

type AuthService struct {
	repo repository.Authorization
}

type TokenClaimsWithId struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum(passwordSalt))
}

func (s *AuthService) CreateUser(user entity.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))

	if err != nil {
		//logrus.Errorf("can not get user from DB: %s", err)
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsWithId{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id,
	})

	return token.SignedString(jwtKey)
}

func (s *AuthService) ParseToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaimsWithId{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*TokenClaimsWithId)

	if !ok {
		return 0, errors.New("token claims are not of type *TokenClaimsWithId")
	}

	return claims.UserId, nil

}
