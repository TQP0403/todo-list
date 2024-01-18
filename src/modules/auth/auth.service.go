package auth

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/auth/dtos"
	"errors"
	"os"
)

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	accessExpire  int
	refreshExpire int
	repo          IAuthRepo
	jwt           IJwtService
}

type IAuthService interface {
	Register(param *dtos.CreateUserDto) (*LoginResponse, error)
	Login(param *dtos.LoginDto) (*LoginResponse, error)
	RefreshToken(token string) (*LoginResponse, error)
}

func NewService(repo *AuthRepo, jwt *JwtService) *AuthService {
	accessExpire := 86400
	if expire := os.Getenv("JWT_ACCESS_EXPIRE"); expire != "" {
		accessExpire = helper.ParseInt(expire)
	}

	refreshExpire := 86400 * 30
	if expire := os.Getenv("JWT_REFRESH_EXPIRE"); expire != "" {
		refreshExpire = helper.ParseInt(expire)
	}

	return &AuthService{jwt: jwt, repo: repo, accessExpire: accessExpire, refreshExpire: refreshExpire}
}

func (service *AuthService) Register(param *dtos.CreateUserDto) (*LoginResponse, error) {
	var user *models.User
	var err error

	if ok, err := service.repo.IsExistUserName(param.Username); err != nil {
		return nil, err
	} else if ok {
		return nil, common.NewConflictError(errors.New("user exits"))
	}

	param.Password, err = helper.BcryptHash(param.Password)
	if err != nil {
		return nil, err
	}

	if user, err = service.repo.CreateUser(param); err != nil {
		return nil, err
	}

	return service.getToken(user.ID)
}

func (service *AuthService) Login(param *dtos.LoginDto) (*LoginResponse, error) {
	var err error
	var user *models.User

	if user, err = service.repo.GetUserByUserName(param.Username); err != nil {
		return nil, err
	}

	// bcrypt compare password
	if err := helper.BcryptCompare(user.Password, param.Password); err != nil {
		return nil, err
	}

	return service.getToken(user.ID)
}

func (service *AuthService) RefreshToken(token string) (*LoginResponse, error) {
	if data, err := service.jwt.JwtVerify(token); err != nil {
		return nil, err
	} else {
		return service.getToken(data.UserId)
	}
}

func (service *AuthService) getToken(id int) (*LoginResponse, error) {
	_, err := service.repo.GetUserById(id)

	if err != nil {
		return nil, err
	}

	accClaim := NewUserCustomClaims(id, service.accessExpire)
	refClaim := NewUserCustomClaims(id, service.refreshExpire)

	res := &LoginResponse{
		AccessToken:  service.jwt.JwtSign(accClaim),
		RefreshToken: service.jwt.JwtSign(refClaim),
	}

	return res, nil
}
