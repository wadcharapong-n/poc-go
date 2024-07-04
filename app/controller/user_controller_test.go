package controller

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserService is a mock implementation of the UserService interface
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) GetAllUser(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) AddUserData(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) GetUserById(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) UpdateUserData(c *gin.Context) {
	m.Called(c)
}

func (m *MockUserService) DeleteUser(c *gin.Context) {
	m.Called(c)
}

func TestUserControllerImpl_GetAllUserData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := UserControllerInit(mockService)

	router := gin.Default()
	router.GET("/users", controller.GetAllUserData)

	req := httptest.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockService.On("GetAllUser", c).Return()

	controller.GetAllUserData(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
}

func TestUserControllerImpl_AddUserData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := UserControllerInit(mockService)

	router := gin.Default()
	router.POST("/users", controller.AddUserData)

	req := httptest.NewRequest("POST", "/users", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	mockService.On("AddUserData", c).Return()

	controller.AddUserData(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
}

func TestUserControllerImpl_GetUserById(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := UserControllerInit(mockService)

	router := gin.Default()
	router.GET("/users/:id", controller.GetUserById)

	req := httptest.NewRequest("GET", "/users/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	mockService.On("GetUserById", c).Return()

	controller.GetUserById(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
}

func TestUserControllerImpl_UpdateUserData(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := UserControllerInit(mockService)

	router := gin.Default()
	router.PUT("/users/:id", controller.UpdateUserData)

	req := httptest.NewRequest("PUT", "/users/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	mockService.On("UpdateUserData", c).Return()

	controller.UpdateUserData(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
}

func TestUserControllerImpl_DeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockUserService)
	controller := UserControllerInit(mockService)

	router := gin.Default()
	router.DELETE("/users/:id", controller.DeleteUser)

	req := httptest.NewRequest("DELETE", "/users/1", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: "1"}}
	c.Request = req

	mockService.On("DeleteUser", c).Return()

	controller.DeleteUser(c)

	mockService.AssertExpectations(t)
	assert.Equal(t, 200, w.Code)
}
