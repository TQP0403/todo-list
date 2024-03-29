package auth

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/helper"
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/auth/dtos"
	"TQP0403/todo-list/src/modules/cache"
	"TQP0403/todo-list/src/modules/jwt"
	"errors"
	"time"

	"gorm.io/gorm"
)

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type AuthService struct {
	accessExpire  int
	refreshExpire int
	jwtService    jwt.IJwtService
	repo          IAuthRepo
	cache         IAuthCache
}

type IAuthService interface {
	Register(param *dtos.RegisterDto) (*LoginResponse, error)
	Login(param *dtos.LoginDto) (*LoginResponse, error)
	RefreshToken(token string) (*LoginResponse, error)
	GetProfile(id int) (*models.User, error)
	UpdateProfile(id int, param *dtos.UpdateProfileDto) error
	ChangePassword(id int, param *dtos.ChangePasswordDto) error
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

	if len(param.DisplayName) == 0 {
		return nil, common.NewBadRequestError(errors.New("displayName is empty"))
	}
	if len(param.Username) == 0 {
		return nil, common.NewBadRequestError(errors.New("username is empty"))
	}
	if len(param.Password) == 0 {
		return nil, common.NewBadRequestError(errors.New("password is empty"))
	}

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, common.NewUnauthorizedError(errors.New("username or password is wrong"))
		}
		return nil, err
	}

	// bcrypt compare password
	if err := helper.BcryptCompare(user.Password, param.Password); err != nil {
		return nil, common.NewUnauthorizedError(errors.New("username or password is wrong"))
	}

	return service.getToken(user.ID)
}

func (service *AuthService) RefreshToken(token string) (*LoginResponse, error) {
	if len(token) == 0 {
		return nil, common.NewBadRequestError(errors.New("refresh token is empty"))
	}

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

	accClaim := jwt.NewUserCustomClaims(id, time.Duration(service.accessExpire)*time.Second)
	refClaim := jwt.NewUserCustomClaims(id, time.Duration(service.refreshExpire)*time.Second)
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

func (service *AuthService) UpdateProfile(id int, param *dtos.UpdateProfileDto) error {
	if param.DisplayName == "" {
		return common.NewBadRequestError(errors.New("display_name is empty"))
	}

	return service.repo.UpdateUser(id, param)
}

func (service *AuthService) ChangePassword(id int, param *dtos.ChangePasswordDto) error {
	// check input params
	if len(param.OldPassword) == 0 {
		return common.NewBadRequestError(errors.New("old password is empty"))
	}
	if len(param.NewPassword) == 0 {
		return common.NewBadRequestError(errors.New("new password is empty"))
	}

	// get user by id & check old password
	if user, err := service.GetProfile(id); err != nil {
		return err
	} else if err = helper.BcryptCompare(user.Password, param.OldPassword); err != nil {
		return common.NewUnauthorizedError(errors.New("wrong password"))
	}

	// hash new password & save password
	if newHashedPassword, err := helper.BcryptHash(param.NewPassword); err != nil {
		return err
	} else {
		return service.repo.ChangePasswordUser(id, newHashedPassword)
	}
}
