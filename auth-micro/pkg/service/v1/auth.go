package v1

import (
	"context"

	"github.com/jinzhu/gorm"
	v1 "github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/api/v1"
	models "github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/models/v1"
)

// AuthService is ...
type AuthService struct {
	db *gorm.DB
}

// NewAuthService returns ...
func NewAuthServer(db *gorm.DB) v1.AuthServer {
	return &AuthService{db: db}
}

// Login is ...
func (a *AuthService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.AuthResponse, error) {
	var user models.User

	userAcc, err := user.FindByUsername(a.db, request.Username)

	if err != nil {
		return &v1.AuthResponse{Status: false}, err
	}

	isValid := userAcc.ValidPassword(request.Password)

	if isValid != nil {
		return &v1.AuthResponse{Status: false}, isValid
	}

	return &v1.AuthResponse{Status: true}, nil
}

// SignUp is ...
func (a *AuthService) SignUp(ctx context.Context, request *v1.UserAuthRequest) (*v1.AuthResponse, error) {

	user := models.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	err := user.Create(a.db)

	if err != nil {
		return &v1.AuthResponse{Status: false}, err
	}

	return &v1.AuthResponse{Status: true}, nil
}
