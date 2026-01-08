package usecase

import (
	"errors"
	"ezkost/internal/domain/entity"
	"ezkost/internal/domain/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase interface {
	Register(user *entity.User, password string) error
	Login(email, password string) (string, *entity.User, error)
	ValidateToken(tokenString string) (uint, string, error)
}

type authUsecase struct {
	userRepo  repository.UserRepository
	jwtSecret string
}

func NewAuthUsecase(userRepo repository.UserRepository, jwtSecret string) AuthUsecase {
	return &authUsecase{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
	}
}

func (u *authUsecase) Register(user *entity.User, password string) error {
	// Check if email already exists
	existing, _ := u.userRepo.FindByEmail(user.Email)
	if existing != nil {
		return errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	return u.userRepo.Create(user)
}

func (u *authUsecase) Login(email, password string) (string, *entity.User, error) {
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}

func (u *authUsecase) ValidateToken(tokenString string) (uint, string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(u.jwtSecret), nil
	})

	if err != nil || !token.Valid {
		return 0, "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, "", errors.New("invalid token claims")
	}

	userID := uint(claims["user_id"].(float64))
	role := claims["role"].(string)

	return userID, role, nil
}
