package task

import (
	"TQP0403/todo-list/src/common"
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/cache"
	"TQP0403/todo-list/src/modules/task/dtos"
	"errors"
)

type ITaskService interface {
	CreateTask(param *dtos.CreateTaskDto) (*models.Task, error)
	GetListTask(userId int, pagination *common.Pagination) ([]*models.Task, error)
	GetTaskById(userId, id int) (*models.Task, error)
	UpdateTask(userId int, param *dtos.UpdateTaskDto) error
	DeleteTask(userId, id int) error
}

type TaskService struct {
	repo  ITaskRepo
	cache ITaskCache
}

func NewService(repo *TaskRepo, cacheService *cache.CacheService) *TaskService {
	return &TaskService{repo: repo, cache: NewTaskCache(cacheService)}
}

func (service *TaskService) CreateTask(param *dtos.CreateTaskDto) (*models.Task, error) {
	if len(param.Title) == 0 {
		return nil, common.NewBadRequestError(errors.New("title is empty"))
	}

	return service.repo.CreateTask(param)
}

func (service *TaskService) GetListTask(userId int, pagination *common.Pagination) ([]*models.Task, error) {
	return service.repo.GetListTask(userId, pagination)
}

func (service *TaskService) GetTaskById(userId, id int) (*models.Task, error) {
	// get cache
	task, err := service.cache.GetCacheTask(id)

	// if cache empty
	if err != nil {
		// get db
		task, err = service.repo.GetTaskById(id)
		if err != nil {
			return nil, err
		}
		// set cache
		service.cache.SetCacheTask(task)
	}

	// check owner
	if task.UserID != userId {
		return nil, common.NewBadRequestError(errors.New("not owner"))
	}

	return task, nil
}

func (service *TaskService) UpdateTask(userId int, param *dtos.UpdateTaskDto) error {
	_, err := service.GetTaskById(userId, param.ID)

	if err != nil {
		return err
	}

	if len(param.Title) == 0 {
		return common.NewBadRequestError(errors.New("title is empty"))
	}

	// delete cache
	service.cache.DelCacheTask(param.ID)

	return service.repo.UpdateTask(param)
}

func (service *TaskService) DeleteTask(userId, id int) error {
	_, err := service.GetTaskById(userId, id)

	if err != nil {
		return err
	}

	// delete cache
	service.cache.DelCacheTask(id)

	return service.repo.DeleteTask(id)
}
