package service

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"time"
	"vault/modules/auth/models"
	"vault/modules/auth/repository"
)

const (
	tokenTTL = 24 * time.Hour * 90
)

type AuthService struct {
	repo       *repository.Repository
	salt       string
	signingKey string
}

type accessTokenClaims struct {
	jwt.StandardClaims
	UserId uuid.UUID `json:"user_id"`
	// polices
}

func (s *AuthService) Login(userLogin models.UserLogin) (string, error) {
	creds, err := s.repo.User.GetUserCreds(userLogin.Login)
	if err != nil {
		logrus.Errorf("modules/auth/service/Login error: %s", err.Error())
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(creds.PasswordHash), []byte(userLogin.Password))
	if err != nil {
		logrus.Errorf("modules/auth/service/Login error: %s", err.Error())
		return "", err
	}

	if creds.MfaEnabled {
		key, err := s.repo.User.GetUserTotpKey(creds.Id)
		if err != nil {
			logrus.Errorf("modules/auth/service/Login error: %s", err.Error())
			return "", err
		}
		valid := totp.Validate(userLogin.Totp, key)
		if !valid {
			logrus.Errorf("modules/auth/service/Login error: %s", "Invalid totp")
			return "", errors.New("invalid totp")
		}
	}

	token, err := s.generateToken(creds.Id)
	if err != nil {
		logrus.Errorf("modules/auth/service/Login error: %s", err.Error())
		return "", err
	}

	return token, nil

}

func (s *AuthService) generateToken(userId uuid.UUID) (string, error) {
	tokenId := uuid.New().String()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, &accessTokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        tokenId,
		},
		userId,
	})

	signedString, err := token.SignedString([]byte(s.signingKey))

	if err != nil {
		logrus.Errorf("modules/auth/service/generateToken error: %s", err.Error())
		return "", err
	}

	return signedString, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))
	return fmt.Sprintf("%x", hash.Sum([]byte(s.salt)))
}

func NewAuthService(repo *repository.Repository) *AuthService {
	salt := viper.GetString("security.salt")
	key := viper.GetString("security.signing_key")
	return &AuthService{repo: repo, salt: salt, signingKey: key}
}
