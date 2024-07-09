package service

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"poc-go/app/domain/dto"
)

type PService interface {
	PA(c *gin.Context)
}

type PServiceImpl struct {
}

func PServiceInit() *PServiceImpl {
	return &PServiceImpl{}
}

func (p PServiceImpl) PA(c *gin.Context) {
	var requestBody dto.PaRequest
	// Bind the JSON to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	marshal, err := json.Marshal(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error Marshal": err.Error()})
		return
	}
	// The URL of the API you want to call
	apiURL := "https://test-fs.free.beeceptor.com/p"

	// Create a new POST request
	request, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(marshal))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
		return
	}
	request.Header.Set("Content-Type", "application/json")

	// Make the POST request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to send request"})
		return
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to read response"})
		return
	}

	// Check for a successful response
	if response.StatusCode != http.StatusOK {
		c.JSON(response.StatusCode, gin.H{"error": "API responded with error", "details": string(body)})
		return
	}
	var responseBody dto.PaResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error Marshal": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responseBody)
}
