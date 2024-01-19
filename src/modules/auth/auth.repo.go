package auth

import (
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/auth/dtos"
	"errors"

	"gorm.io/gorm"
)

type IAuthRepo interface {
	IsExistUserName(username string) (bool, error)
	GetUserByUserName(username string) (*models.User, error)
	GetUserById(id int) (*models.User, error)
	CreateUser(param *dtos.RegisterDto) (*models.User, error)
	UpdateUser(id int, param *dtos.UpdateProfileDto) error
	ChangePasswordUser(id int, newPassword string) error
}

type AuthRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (repo *AuthRepo) IsExistUserName(username string) (bool, error) {
	if _, err := repo.GetUserByUserName(username); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (repo *AuthRepo) GetUserByUserName(username string) (*models.User, error) {
	user := &models.User{}

	err := repo.db.Model(&models.User{}).
		Where("username = ?", username).
		First(user).
		Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *AuthRepo) GetUserById(id int) (*models.User, error) {
	user := &models.User{}

	err := repo.db.Model(&models.User{}).
		Where("id = ?", id).
		First(user).
		Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *AuthRepo) CreateUser(param *dtos.RegisterDto) (*models.User, error) {
	user := &models.User{
		DisplayName: param.DisplayName,
		Username:    param.Username,
		Password:    param.Password,
	}

	if err := repo.db.Model(&models.User{}).Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *AuthRepo) UpdateUser(id int, param *dtos.UpdateProfileDto) error {
	err := repo.db.Model(&models.User{}).Where("id = ?", id).Updates(param).Error
	if err != nil {
		return err
	}
	return nil
}

func (repo *AuthRepo) ChangePasswordUser(id int, newPassword string) error {
	err := repo.db.Model(&models.User{}).Where("id = ?", id).Update("password", newPassword).Error
	if err != nil {
		return err
	}
	return nil
}
