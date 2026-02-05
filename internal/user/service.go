package user

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var SecretKey = []byte("super-secret-key") // In prod, use os.Getenv("SECRET_KEY")

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
	//Normalize email
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))

	// check if email already exists
	exists, err := s.Repo.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email already exists")
	}
	// Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil {
		return err
	}
	user.Password = string(hashed)

	return s.Repo.Create(user)
}

func (s *Service) Login(email, password string) (*User, string, error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, "", errors.New("invalid credentials")
	}

	now := time.Now()
	user.LastLogin = &now
	s.Repo.Update(user)

	tokenStr, err := generateToken(user.ID)
	if err != nil {
		return nil, "", err
	}

	return user, tokenStr, nil
}

// Helper function (if not already separated)
func generateToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString(SecretKey)
}
