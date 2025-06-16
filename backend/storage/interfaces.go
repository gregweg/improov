package storage

import "improov/models"

type DBInterface interface {
	GetOrCreateUser(userID string) (*models.User, error)
	GetTasksByCategory(category string) ([]models.Task, error)
	SaveUser(user *models.User) error
	CreateCompletedTask(userID, taskID string) error
	GetCompletedTasks(userID string) ([]models.CompletedTask, error)
	GetAllCategories(categories *[]models.Category) error
	GetUserByID(userID string) (*models.User, error)
}
