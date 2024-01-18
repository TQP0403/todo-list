package auth

import (
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/auth/dtos"

	"gorm.io/gorm"
)

type IAuthRepo interface {
	IsExistUserName(username string) (bool, error)
	GetUserByUserName(username string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	CreateUser(param *dtos.CreateUserDto) (*models.User, error)
}

type AuthRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (repo *AuthRepo) IsExistUserName(username string) (bool, error) {
	if _, err := repo.GetUserByUserName(username); err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *AuthRepo) GetUserByUserName(username string) (*models.User, error) {
	user := &models.User{}

	if err := repo.db.Model(&models.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *AuthRepo) GetUserById(id int) (*models.User, error) {
	user := &models.User{}

	if err := repo.db.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *AuthRepo) CreateUser(param *dtos.CreateUserDto) (*models.User, error) {
	user := &models.User{
		DisplayName: param.DisplayName,
		Username:    param.Username,
		Password:    param.Password,
	}

	if err := repo.db.Model(&models.User{}).Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
