package usecase

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/MaryneZa/tafins-backend/entity"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	SignUp(user entity.User) error
	LogIn(user entity.User) (string, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserUseCase {
	return &UserService{repo: repo}
}

func (us *UserService) SignUp(user entity.User) error {
	hashPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashPassword
	return us.repo.Create(user)
}

func (us *UserService) LogIn(user entity.User) (string, error) {
	userDetail, err := us.repo.Get(user)
	if err != nil {
		return "", err
	}
	if err := us.repo.Find(user.Email); err != nil {
		return "", err
	}
	if err := CheckPasswordHash(user.Password, userDetail.Password); !err {
		return "", fmt.Errorf("wrong password")
	}
	token, err := CreateToken(userDetail)
	if err != nil {
		return "", err
	}
	return token, nil
}

func CreateToken(user entity.User) (string, error) {
	secretKey := os.Getenv("JWT_SECRETKEY")
	log.Println("secretKey: ", secretKey, "user.ID: ", user.ID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user_id": user.ID,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
