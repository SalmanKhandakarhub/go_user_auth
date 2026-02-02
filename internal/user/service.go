package user

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var Secretkey = []byte("super-secret-key") // In prod, use os.Getenv("SECRET_KEY")

type Service struct {
	Repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{
		Repo: repo,
	}
}

// --- Logic Implementation ---

func (s *Service) Register(user *User) error {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashed)
	return s.Repo.Create(user)
}

func (s *Service) Login(email, password string) (string, error) {
	user, err := s.Repo.FindByEmail(email)

	if err != nil {
		return "", errors.New("Invalid Credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return "", errors.New("Invalid Credentials")
	}

	now := time.Now()
	user.LastLogin = &now
	s.Repo.Update(user)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(Secretkey)
}
