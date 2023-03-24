package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/deevins/todo-restAPI/internal/entities"
	"github.com/deevins/todo-restAPI/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	uuid "github.com/satori/go.uuid"
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

// TokenClaimsWithId Structure of token with user id
type TokenClaimsWithId struct {
	jwt.RegisteredClaims
	UserId uuid.UUID `json:"user_id"`
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

func (s *AuthService) CreateUser(user entity.User) (uuid.UUID, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	// token create new JWT token and sign it.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenClaimsWithId{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},

		UserId: user.Id,
	})

	return token.SignedString(jwtKey)
}

func (s *AuthService) ParseToken(accessToken string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(accessToken, &TokenClaimsWithId{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		return Nil, err
	}

	claims, ok := token.Claims.(*TokenClaimsWithId)

	if !ok {
		return Nil, errors.New("token claims are not of type *TokenClaimsWithId")
	}

	return claims.UserId, nil

}
