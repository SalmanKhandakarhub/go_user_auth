package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint       `gorm:"primaryKey" json:"id"`
	FullName       string     `gorm:"size:64;index" json:"full_name"`
	Email          string     `gorm:"size:64;uniqueIndex" json:"email"`
	Password       string     `gorm:"size:255" json:"-"`
	ContactNo      *string    `gorm:"size:20;index" json:"contact_no"`
	ProfileImage   *string    `gorm:"size:255;index" json:"profile_image"`
	Address        *string    `gorm:"size:255;index" json:"address"`
	IsVerified     bool       `gorm:"default:false" json:"is_verified"`
	LastLogin      *time.Time `json:"last_login"`
	SocialProvider *string    `gorm:"size:64;index" json:"social_provider"`
	IsAdmin        bool       `gorm:"default:false" json:"is_admin"`

	// Standard GORM timestamp fields
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
