package task

import (
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/cache"
	"fmt"
	"time"
)

type ITaskCache interface {
	GetCacheTask(id int) (*models.Task, error)
	SetCacheTask(user *models.Task) error
	DelCacheTask(id int) error

	GetCacheListTasks(userId int) ([]models.Task, error)
	SetCacheListTasks(userId int, tasks []models.Task) error
	DelCacheListTasks(userId int) error
}

type TaskCache struct {
	cacheService cache.ICacheService
}

func NewTaskCache(cacheService *cache.CacheService) *TaskCache {
	return &TaskCache{cacheService: cacheService}
}

func (service *TaskCache) GetCacheTask(id int) (*models.Task, error) {
	task := &models.Task{}
	key := fmt.Sprintf("task:%d", id)
	if err := service.cacheService.Get(key, &task); err != nil {
		return nil, err
	}

	return task, nil
}

func (service *TaskCache) SetCacheTask(task *models.Task) error {
	key := fmt.Sprintf("task:%d", task.ID)
	ttl := time.Minute * 5
	if err := service.cacheService.Set(key, &task, ttl); err != nil {
		return err
	}

	return nil
}

func (service *TaskCache) DelCacheTask(id int) error {
	key := fmt.Sprintf("task:%d", id)
	if err := service.cacheService.Del(key); err != nil {
		return err
	}

	return nil
}

// TODO: GetCacheListTasks
func (service *TaskCache) GetCacheListTasks(userId int) ([]models.Task, error) {
	tasks := []models.Task{}
	// key := fmt.Sprintf("list:task:%d", userId)

	return tasks, nil
}

// TODO: SetCacheListTasks
func (service *TaskCache) SetCacheListTasks(userId int, tasks []models.Task) error {
	key := fmt.Sprintf("list:task:%d", userId)
	ttl := time.Minute * 5
	if err := service.cacheService.Set(key, tasks, ttl); err != nil {
		return err
	}
	return nil
}

// TODO: DelCacheListTasks
func (service *TaskCache) DelCacheListTasks(userId int) error {
	key := fmt.Sprintf("list:task:%d", userId)
	if err := service.cacheService.Del(key); err != nil {
		return err
	}
	return nil
}
