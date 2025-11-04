package usecase

import (
	"errors"
	"strings"
	"time"

	"github.com/B1gdawg0/real_course_review_backend/internal/dtos"
	"github.com/B1gdawg0/real_course_review_backend/internal/model"
	"github.com/B1gdawg0/real_course_review_backend/internal/repository"
	"github.com/golang-jwt/jwt/v5"
)

type authUsecase struct {
	repo      repository.AuthRepository
	jwtSecret string
}

func NewAuthUsecase(repo repository.AuthRepository, jwtSecret string) AuthUsecase {
	return &authUsecase{repo: repo, jwtSecret: jwtSecret}
}

func (u *authUsecase) LoginOrRegister(req *dtos.AuthRequest) (*dtos.AuthResponse, error) {
	if !strings.HasSuffix(req.Email, "@ku.th") {
		return nil, errors.New("unauthorized domain")
	}

	user, err := u.repo.GetUserByEmail(req.Email)
	if err != nil {
		user = &model.User{
			Name:  req.Name,
			Email: req.Email,
			Role:  "USER",
		}
		if err := u.repo.CreateUser(user); err != nil {
			return nil, err
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":   user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.jwtSecret))
	if err != nil {
		return nil, err
	}

	resp := &dtos.AuthResponse{
		Token: tokenString,
		User:  dtos.UserShortResponse{ID: user.ID, Name: user.Name},
	}
	return resp, nil
}
