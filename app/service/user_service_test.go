package service

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"poc-go/app/domain/dao"
	"poc-go/app/domain/dto"
	"strings"
	"testing"
)

// MockUserRepository is a mock implementation of the UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUserById(userID int) (dao.User, error) {
	args := m.Called(userID)
	return args.Get(0).(dao.User), args.Error(1)
}

func (m *MockUserRepository) Save(user *dao.User) (dao.User, error) {
	args := m.Called(user)
	return args.Get(0).(dao.User), args.Error(1)
}

func (m *MockUserRepository) FindAllUser() ([]dao.User, error) {
	args := m.Called()
	return args.Get(0).([]dao.User), args.Error(1)
}

func (m *MockUserRepository) DeleteUserById(userID int) error {
	args := m.Called(userID)
	return args.Error(0)
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestUserServiceImpl_UpdateUserData(t *testing.T) {
	router := setupRouter()
	mockRepo := new(MockUserRepository)
	service := UserServiceInit(mockRepo)
	user := dao.User{ID: 1, RoleID: 1, Email: "test@test.com", Password: "password", Status: 1}
	updatedUser := user
	updatedUser.RoleID = 2
	updatedUser.Email = "newemail@test.com"

	router.PUT("/users/:userID", service.UpdateUserData)

	mockRepo.On("FindUserById", 1).Return(user, nil)
	mockRepo.On("Save", mock.Anything).Return(updatedUser, nil)

	w := httptest.NewRecorder()
	body := `{"roleID": 2, "email": "newemail@test.com", "status": 1}`
	req, _ := http.NewRequest("PUT", "/users/1", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	var userResponse dto.ApiResponse[dao.User]
	_ = json.Unmarshal(w.Body.Bytes(), &userResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, updatedUser.Email, userResponse.Data.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserServiceImpl_GetUserById(t *testing.T) {
	router := setupRouter()
	mockRepo := new(MockUserRepository)
	service := UserServiceInit(mockRepo)
	user := dao.User{ID: 1, RoleID: 1, Email: "test@test.com", Password: "password", Status: 1}

	router.GET("/users/:userID", service.GetUserById)

	mockRepo.On("FindUserById", 1).Return(user, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users/1", nil)

	router.ServeHTTP(w, req)

	var userResponse dto.ApiResponse[dao.User]
	_ = json.Unmarshal(w.Body.Bytes(), &userResponse)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, user.Email, userResponse.Data.Email)
	mockRepo.AssertExpectations(t)
}

func TestUserServiceImpl_AddUserData(t *testing.T) {
	router := setupRouter()
	mockRepo := new(MockUserRepository)
	service := UserServiceInit(mockRepo)
	user := dao.User{RoleID: 1, Email: "test@test.com", Password: "password", Status: 1}

	router.POST("/users", service.AddUserData)

	mockRepo.On("Save", mock.Anything).Return(user, nil)

	w := httptest.NewRecorder()
	body := `{"roleID": 1, "email": "test@test.com", "password": "password", "status": 1}`
	req, _ := http.NewRequest("POST", "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestUserServiceImpl_GetAllUser(t *testing.T) {
	router := setupRouter()
	mockRepo := new(MockUserRepository)
	service := UserServiceInit(mockRepo)
	users := []dao.User{
		{ID: 1, RoleID: 1, Email: "test1@test.com", Password: "password", Status: 1},
		{ID: 2, RoleID: 1, Email: "test2@test.com", Password: "password", Status: 1},
	}

	router.GET("/users", service.GetAllUser)

	mockRepo.On("FindAllUser").Return(users, nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}

func TestUserServiceImpl_DeleteUser(t *testing.T) {
	router := setupRouter()
	mockRepo := new(MockUserRepository)
	service := UserServiceInit(mockRepo)

	router.DELETE("/users/:userID", service.DeleteUser)

	mockRepo.On("DeleteUserById", 1).Return(nil)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/users/1", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.AssertExpectations(t)
}
