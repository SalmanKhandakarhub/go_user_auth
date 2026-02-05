package user

import (
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		DB: db,
	}
}

func (r *Repository) Create(user *User) error {
	return r.DB.Create(user).Error
}

func (r *Repository) Update(user *User) error {
	return r.DB.Save(user).Error
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User
	err := r.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *Repository) ExistsByEmail(email string) (bool, error) {
	var count int64
	err := r.DB.Model(&User{}).
		Where("email = ?", email).
		Count(&count).Error
	return count > 0, err
}
