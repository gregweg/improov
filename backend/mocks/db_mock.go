package mocks

import (
	"improov/models"

	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) GetOrCreateUser(userID string) (*models.User, error) {
	args := m.Called(userID)
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockDB) GetTasksByCategory(category string) ([]models.Task, error) {
	args := m.Called(category)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockDB) SaveUser(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockDB) CreateCompletedTask(userID, taskID string) error {
	args := m.Called(userID, taskID)
	return args.Error(0)
}

func (m *MockDB) GetCompletedTasks(userID string) ([]models.CompletedTask, error) {
	args := m.Called(userID)
	return args.Get(0).([]models.CompletedTask), args.Error(1)
}

func (m *MockDB) GetAllCategories(out *[]models.Category) error {
	args := m.Called(out)
	if res := args.Get(0); res != nil {
		*out = res.([]models.Category)
	}
	return args.Error(1)
}

func (m *MockDB) GetUserByID(userID string) (*models.User, error) {
	args := m.Called(userID)
	user := args.Get(0)
	if user == nil {
		return nil, args.Error(1)
	}
	return user.(*models.User), args.Error(1)
}
