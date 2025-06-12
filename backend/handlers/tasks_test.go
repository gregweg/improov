package handlers

import (
	"encoding/json"
	"improov/auth"
	"improov/middleware"
	"improov/mocks"
	"improov/models"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getAuthToken(t *testing.T) string {
	form := url.Values{}
	form.Add("username", "admin")
	form.Add("password", "password123")

	resp, err := http.PostForm("http://localhost:8080/api/login", form)
	if err != nil {
		t.Fatalf("Failed to generate test token: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		t.Fatalf("Failed to decode login response: %v", err)
	}

	token, ok := result["token"]
	if !ok {
		t.Fatalf("No token in login response")
	}

	return token
}
func wrapWithJWTAuth(handler http.HandlerFunc) http.HandlerFunc {
	return middleware.JWTAuthMiddleware(handler).ServeHTTP
}

func getTestToken(t *testing.T) string {
	token, err := auth.GenerateJWT("tester")
	if err != nil {
		t.Fatalf("Failed to generate test token: %v", err)
	}
	return token
}

func TestSuggestTaskSuccess(t *testing.T) {
	mockDB := new(mocks.MockDB)
	handler := TaskHandler{DB: mockDB}

	mockTasks := []models.Task{
		{ID: "fit-001", Category: "fitness", Description: "Do 10 push-ups"},
	}

	mockDB.On("GetTasksByCategory", "fitness").Return(mockTasks, nil)

	req := httptest.NewRequest("GET", "/api/tasks/suggest?category=fitness&userId=tester", nil)
	req.Header.Set("Authorization", "Bearer "+getTestToken(t))
	rr := httptest.NewRecorder()

	wrapWithJWTAuth(handler.SuggestTask)(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var task models.Task
	body := rr.Body.Bytes()
	t.Logf("Response body: %s", string(body))
	err := json.Unmarshal(rr.Body.Bytes(), &task)
	assert.NoError(t, err)
	assert.Equal(t, "fitness", task.Category)

	mockDB.AssertExpectations(t)
}

func TestSuggestTaskMissingParams(t *testing.T) {
	mockDB := new(mocks.MockDB)
	handler := TaskHandler{DB: mockDB}

	req := httptest.NewRequest("GET", "/api/tasks/suggest?category=fitness", nil)
	rr := httptest.NewRecorder()

	handler.SuggestTask(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
}

func TestCompleteTaskSuccess(t *testing.T) {
	mockDB := new(mocks.MockDB)
	handler := TaskHandler{DB: mockDB}

	// Consistent test user
	userID := "admin"

	user := &models.User{
		ID:          userID,
		Fitness:     0,
		Learning:    0,
		Mindfulness: 0,
	}

	// Expect these mock calls for user "tester"
	mockDB.On("GetOrCreateUser", userID).Return(user, nil)
	mockDB.On("SaveUser", user).Return(nil)
	mockDB.On("CreateCompletedTask", userID, "fit-001").Return(nil)

	body := `{"user_id":"admin","task_id":"fit-001","category":"fitness"}`

	req := httptest.NewRequest("POST", "/api/tasks/complete", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getTestToken(t)) // token for "tester"

	rr := httptest.NewRecorder()
	wrapWithJWTAuth(handler.CompleteTask)(rr, req)

	respBody := rr.Body.Bytes()
	t.Logf("Response body: %s", string(respBody))

	var updated models.User
	err := json.Unmarshal(respBody, &updated)
	assert.NoError(t, err)

	assert.Equal(t, 10, updated.Fitness)
	assert.Equal(t, 0, updated.Learning)
	assert.Equal(t, 0, updated.Mindfulness)

	mockDB.AssertExpectations(t)
}

func TestGetCompletedTasks(t *testing.T) {
	mockDB := new(mocks.MockDB)
	handler := TaskHandler{DB: mockDB}

	userID := "tester"
	// Simulate a completed task with task preloaded
	mockCompleted := []models.CompletedTask{
		{
			TaskID: "fit-001",
			Task: models.Task{
				ID:          "fit-001",
				Category:    "Fitness",
				Description: "Do 10 push-ups",
			},
		},
	}

	mockDB.On("GetCompletedTasks", userID).Return(mockCompleted, nil)

	req := httptest.NewRequest("GET", "/api/tasks/completed?userId=tester", nil)
	req.Header.Set("Authorization", "Bearer "+getTestToken(t))
	rr := httptest.NewRecorder()

	wrapWithJWTAuth(handler.GetCompletedTasks)(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var result []models.CompletedTask
	err := json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NoError(t, err)

	assert.Len(t, result, 1)
	assert.Equal(t, "Do 10 push-ups", result[0].Task.Description)
	assert.Equal(t, "fit-001", result[0].Task.ID)

	mockDB.AssertExpectations(t)
}
