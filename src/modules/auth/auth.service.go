package auth

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/auth/dtos"
	"TQP0403/todo-list/src/modules/cache"
	"TQP0403/todo-list/src/modules/jwt"
	"errors"
)

type LoginResponse struct {
	AccessToken  string
	RefreshToken string
}

type AuthService struct {
	accessExpire  int
	refreshExpire int
	jwtService    jwt.IJwtService
	repo          IAuthRepo
	cache         IAuthCache
}

type IAuthService interface {
	GetJwtService() jwt.IJwtService
	Register(param *dtos.RegisterDto) (*LoginResponse, error)
	Login(param *dtos.LoginDto) (*LoginResponse, error)
	RefreshToken(token string) (*LoginResponse, error)
	GetProfile(id int) (*models.User, error)
}

func (service *AuthService) GetJwtService() jwt.IJwtService {
	return service.jwtService
}

func NewService(repo *AuthRepo, jwtService *jwt.JwtService, cacheService *cache.CacheService) *AuthService {
	accessExpire := helper.ParseInt(helper.GetDefaultEnv("JWT_ACCESS_EXPIRE", "86400"))
	refreshExpire := helper.ParseInt(helper.GetDefaultEnv("JWT_REFRESH_EXPIRE", "2592000"))
	cache := NewAuthCache(cacheService)

	return &AuthService{
		repo:          repo,
		cache:         cache,
		jwtService:    jwtService,
		accessExpire:  accessExpire,
		refreshExpire: refreshExpire,
	}
}

func (service *AuthService) Register(param *dtos.RegisterDto) (*LoginResponse, error) {
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

	accClaim := jwt.NewUserCustomClaims(id, service.accessExpire)
	refClaim := jwt.NewUserCustomClaims(id, service.refreshExpire)
	res := &LoginResponse{
		AccessToken:  service.jwtService.JwtSign(accClaim),
		RefreshToken: service.jwtService.JwtSign(refClaim),
	}

	return res, nil
}

func (service *AuthService) GetProfile(id int) (*models.User, error) {
	// get cache
	user, err := service.cache.GetCacheUser(id)

	// if cache empty
	if err != nil {
		// get db
		user, err = service.repo.GetUserById(id)
		if err != nil {
			return nil, err
		}
		// set cache
		service.cache.SetCacheUser(user)
	}

	return user, nil
}
