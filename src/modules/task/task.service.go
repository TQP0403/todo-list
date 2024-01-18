package task

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/task/dtos"
	"errors"
)

type ITaskService interface {
	CreateTask(param *dtos.CreateTaskDto) (*models.Task, error)
	GetListTask(userId int, pagination *common.Pagination) ([]models.Task, error)
	GetTaskById(userId, id int) (*models.Task, error)
	UpdateTask(userId int, param *dtos.UpdateTaskDto) (*models.Task, error)
	DeleteTask(userId, id int) error
}

type TaskService struct {
	repo ITaskRepo
}

func NewService(repo ITaskRepo) *TaskService {
	return &TaskService{repo: repo}
}

func (service *TaskService) CreateTask(param *dtos.CreateTaskDto) (*models.Task, error) {
	return service.repo.CreateTask(param)
}

func (service *TaskService) GetListTask(userId int, pagination *common.Pagination) ([]models.Task, error) {
	return service.repo.GetListTask(userId, pagination)
}

func (service *TaskService) GetTaskById(userId, id int) (*models.Task, error) {
	task, err := service.repo.GetTaskById(id)

	if err != nil {
		return nil, err
	}

	if task.UserID != userId {
		return nil, errors.New("not owner")
	}

	return task, nil
}

func (service *TaskService) UpdateTask(userId int, param *dtos.UpdateTaskDto) (*models.Task, error) {
	_, err := service.GetTaskById(userId, param.ID)

	if err != nil {
		return nil, err
	}

	return service.repo.UpdateTask(param)
}

func (service *TaskService) DeleteTask(userId, id int) error {
	_, err := service.GetTaskById(userId, id)

	if err != nil {
		return err
	}

	return service.repo.DeleteTask(id)
}
