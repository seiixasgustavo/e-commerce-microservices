package v1

import (
	"context"

	"github.com/jinzhu/gorm"
	v1 "github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/api/v1"
	models "github.com/seiixasgustavo/e-commerce-microservices/auth-micro/pkg/models/v1"
)

// UserService is ...
type UserService struct {
	db *gorm.DB
}

// NewUserServer is ...
func NewUserServer(db *gorm.DB) v1.UserServer {
	db.AutoMigrate(&models.User{})
	return &UserService{db: db}
}

// Create is ...
func (u *UserService) Create(ctx context.Context, request *v1.UserRequest) (*v1.Response, error) {
	user := models.User{
		Username: request.User.GetUsername(),
		Email:    request.User.GetEmail(),
		Password: request.User.GetPassword(),
	}

	err := user.Create(u.db)

	if err != nil {
		return &v1.Response{Status: false}, err
	}

	return &v1.Response{Status: true}, nil
}

// Update is ...
func (u *UserService) Update(ctx context.Context, request *v1.UserIdRequest) (*v1.Response, error) {
	user := models.User{
		Username: request.User.GetUsername(),
		Email:    request.User.GetEmail(),
		Password: request.User.GetPassword(),
	}

	err := user.Update(u.db, uint(request.ID))

	if err != nil {
		return &v1.Response{Status: false}, err
	}

	return &v1.Response{Status: true}, nil
}

//Delete is ...
func (u *UserService) Delete(ctx context.Context, request *v1.IdRequest) (*v1.Response, error) {
	var user models.User

	err := user.Delete(u.db, uint(request.ID))

	if err != nil {
		return &v1.Response{Status: false}, err
	}

	return &v1.Response{Status: true}, nil
}

// ChangePassword is ...
func (u *UserService) ChangePassword(ctx context.Context, request *v1.PasswordRequest) (*v1.Response, error) {
	user := models.User{Password: request.GetPassword()}

	err := user.ChangePassword(u.db, uint(request.GetID()))

	if err != nil {
		return &v1.Response{Status: false}, err
	}

	return &v1.Response{Status: true}, nil
}

// FindByPk is ...
func (u *UserService) FindByPk(ctx context.Context, request *v1.IdRequest) (*v1.UserResponse, error) {
	user := models.User{}

	usr, err := user.FindByPk(u.db, uint(request.GetID()))

	if err != nil {
		return &v1.UserResponse{Status: false}, err
	}

	return &v1.UserResponse{
		Status: true,
		User: &v1.UserStruct{
			ID:       uint64(usr.ID),
			Username: usr.Username,
			Email:    usr.Email,
			Password: "",
		},
	}, nil
}

// FindByUsername is ...
func (u *UserService) FindByUsername(ctx context.Context, request *v1.UsernameRequest) (*v1.UserResponse, error) {
	user := models.User{}

	usr, err := user.FindByUsername(u.db, request.GetUsername())

	if err != nil {
		return &v1.UserResponse{Status: false}, err
	}

	return &v1.UserResponse{
		Status: true,
		User: &v1.UserStruct{
			ID:       uint64(usr.ID),
			Username: usr.Username,
			Email:    usr.Email,
			Password: "",
		},
	}, nil
}
