package auth

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/auth/dtos"
	"TQP0403/todo-list/src/modules/cache"
	"errors"
	"os"

	"github.com/redis/go-redis/v9"
)

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	accessExpire  int
	refreshExpire int
	repo          IAuthRepo
	jwtService    IJwtService
	authCache     IAuthCache
}

type IAuthService interface {
	GetJwtService() IJwtService
	Register(param *dtos.CreateUserDto) (*LoginResponse, error)
	Login(param *dtos.LoginDto) (*LoginResponse, error)
	RefreshToken(token string) (*LoginResponse, error)
	GetProfile(id int) (*models.User, error)
}

func (service *AuthService) GetJwtService() IJwtService {
	return service.jwtService
}

func NewService(repo *AuthRepo, jwtService *JwtService, cacheService *cache.CacheService) *AuthService {
	accessExpire := 86400
	if expire := os.Getenv("JWT_ACCESS_EXPIRE"); expire != "" {
		accessExpire = helper.ParseInt(expire)
	}
	refreshExpire := 86400 * 30
	if expire := os.Getenv("JWT_REFRESH_EXPIRE"); expire != "" {
		refreshExpire = helper.ParseInt(expire)
	}

	authCache := NewAuthCache(cacheService)
	return &AuthService{
		repo:          repo,
		authCache:     authCache,
		jwtService:    jwtService,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
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
	if data, err := service.jwtService.JwtVerify(token); err != nil {
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
		AccessToken:  service.jwtService.JwtSign(accClaim),
		RefreshToken: service.jwtService.JwtSign(refClaim),
	}

	return res, nil
}

func (service *AuthService) GetProfile(id int) (*models.User, error) {
	user, err := service.authCache.GetCacheUser(id)

	if errors.Is(err, redis.Nil) {
		user, err = service.repo.GetUserById(id)
	}
	if err != nil {
		return nil, err
	}
	if err = service.authCache.SetCacheUser(user); err != nil {
		return nil, err
	}

	return user, nil
}
