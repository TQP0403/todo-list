package task

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/task/dtos"

	"gorm.io/gorm"
)

type ITaskRepo interface {
	CreateTask(param *dtos.CreateTaskDto) (*models.Task, error)
	GetListTask(userId int, pagination *common.Pagination) ([]*models.Task, error)
	GetTaskById(id int) (*models.Task, error)
	UpdateTask(param *dtos.UpdateTaskDto) error
	DeleteTask(id int) error
}

type TaskRepo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *TaskRepo {
	return &TaskRepo{db: db}
}

func (repo *TaskRepo) CreateTask(param *dtos.CreateTaskDto) (*models.Task, error) {
	task := &models.Task{
		UserID:  param.UserID,
		Title:   param.Title,
		Content: param.Content,
		Status:  param.Status,
	}

	if err := repo.db.Model(&models.Task{}).Create(&task).Error; err != nil {
		return nil, err
	}

	return task, nil
}

func (repo *TaskRepo) GetListTask(userId int, pagination *common.Pagination) ([]*models.Task, error) {
	var tasks []*models.Task

	err := repo.db.Model(&models.Task{}).
		Where("user_id = ?", userId).
		Count(&pagination.Total).
		Offset((pagination.Page - 1) * pagination.PageSize).
		Limit(pagination.PageSize).
		Order("id desc").
		Find(&tasks).Error

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repo *TaskRepo) GetTaskById(id int) (*models.Task, error) {
	var task models.Task

	if err := repo.db.Model(&models.Task{}).Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}

	return &task, nil
}

func (repo *TaskRepo) UpdateTask(param *dtos.UpdateTaskDto) error {
	if err := repo.db.Model(&models.Task{}).Where("id = ?", param.ID).Updates(param).Error; err != nil {
		return err
	}

	return nil
}

func (repo *TaskRepo) DeleteTask(id int) error {
	if err := repo.db.Model(&models.Task{}).Delete(&models.Task{ID: id}).Error; err != nil {
		return err
	}

	return nil
}
