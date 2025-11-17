package repository

import (
	m "github.com/B1gdawg0/real_course_review_backend/internal/model"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) GetUserByEmail(email string) (*m.User, error) {
	var user m.User
	if err := r.db.Where("email_address = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authRepository) VerifyUserById(id string) (bool, error) {
    var count int64
    err := r.db.Model(&m.User{}).Where("id = ?", id).Count(&count).Error
    if err != nil {
        return false, err
    }
    return count > 0, nil
}

func (r *authRepository) CreateUser(user *m.User) error {
	return r.db.Create(user).Error
}
