package task

import (
	"TQP0403/todo-list/src/models"
	"TQP0403/todo-list/src/modules/cache"
	"fmt"
	"time"
)

type ITaskCache interface {
	GetCacheTask(id int) (*models.Task, error)
	SetCacheTask(task *models.Task) error
	DelCacheTask(id int) error
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
	if err := service.cacheService.Get(key, task); err != nil {
		return nil, err
	}

	return task, nil
}

func (service *TaskCache) SetCacheTask(task *models.Task) error {
	ttl := time.Minute * 5
	key := fmt.Sprintf("task:%d", task.ID)
	if err := service.cacheService.Set(key, task, ttl); err != nil {
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
