package services

import (
	"errors"
	"shopping-website/models"
	"time"

	"github.com/dgrijalva/jwt-go" // 用於生成 JWT
	"golang.org/x/crypto/bcrypt"  // 用於密碼加密
)

type AuthService struct {
	userModel *models.UserModel
	secretKey []byte // JWT 密鑰
}

func NewAuthService(userModel *models.UserModel, secretKey string) *AuthService {
	return &AuthService{
		userModel: userModel,
		secretKey: []byte(secretKey),
	}
}

// Register handles user registration.
func (s *AuthService) Register(user *models.User) error {
	existingUser, err := s.userModel.FindUser(user.Email)
	if err != nil {
		return err // 返回查找用戶時的錯誤
	}
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// 密碼加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	return s.userModel.CreateUser(user)
}

// Login handles user login and returns a JWT token.
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.userModel.FindUser(email)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	// 驗證密碼
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	// 生成 JWT
	token, err := s.generateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

// generateToken generates a JWT token for the user.
func (s *AuthService) generateToken(user *models.User) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour) // 設定過期時間為 24 小時

	claims := &jwt.StandardClaims{
		Issuer:    user.Email,
		ExpiresAt: expirationTime.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
