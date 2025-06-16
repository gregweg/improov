package storage

import (
	"improov/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBImpl struct{}

func Init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("improov.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto migrate all models
	DB.AutoMigrate(&models.User{}, &models.Task{}, &models.CompletedTask{}, &models.Category{})
}

// DBInterface Implementation
func (db *DBImpl) GetTasksByCategory(category string) ([]models.Task, error) {
	var tasks []models.Task
	result := DB.Where("lower(category) = ?", category).Find(&tasks)
	return tasks, result.Error
}

func (db *DBImpl) GetOrCreateUser(userID string) (*models.User, error) {
	var user models.User
	result := DB.FirstOrCreate(&user, models.User{ID: userID})
	return &user, result.Error
}

func (db *DBImpl) SaveUser(user *models.User) error {
	return DB.Save(user).Error
}

func (db *DBImpl) CreateCompletedTask(userID, taskID string) error {
	return DB.Create(&models.CompletedTask{UserID: userID, TaskID: taskID}).Error
}

func (db *DBImpl) GetCompletedTasks(userID string) ([]models.CompletedTask, error) {
	var completedTasks []models.CompletedTask
	result := DB.Preload("Task").Where("user_id = ?", userID).Find(&completedTasks)
	return completedTasks, result.Error
}

func (db *DBImpl) GetAllCategories(categories *[]models.Category) error {
	return DB.Find(categories).Error
}

func (db *DBImpl) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := DB.First(&user, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
