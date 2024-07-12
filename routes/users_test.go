package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"events.com/models"
	"events.com/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestLoginUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	// Create a test router
	router := gin.New()
	router.POST("/login", routes.LoginUser)

	// Create a test user
	user := models.User{
		Email:    "test@example.com",
		Password: "password123",
	}

	// Save the test user
	err := user.Save()
	assert.NoError(t, err)

	// Marshal the user into JSON
	payload, err := json.Marshal(user)
	assert.NoError(t, err)

	// Create a request with the JSON payload
	req, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	// Set the request content type to JSON
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	res := httptest.NewRecorder()

	// Perform the request
	router.ServeHTTP(res, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, res.Code)

	// Decode the response JSON
	var response map[string]string
	err = json.Unmarshal(res.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Check the response token
	assert.NotEmpty(t, response["token"])

	// Clean up the test user
	// err = user.Delete()
	assert.NoError(t, err)
}
