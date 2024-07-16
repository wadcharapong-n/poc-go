package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"poc-go/app/domain/dto"
)

type PService interface {
	PA(c *gin.Context)
	BR(c *gin.Context)
	DO(c *gin.Context)
	SavePac(c *gin.Context)
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

func (p PServiceImpl) BR(c *gin.Context) {

	// Parse the query parameters
	queryValues := c.Request.URL.Query()

	// Access specific query parameter
	branchID := queryValues.Get("branch_id")
	requestID := queryValues.Get("requestID")
	timezone := queryValues.Get("timezone")
	clientDetail := queryValues.Get("client_detail")
	wnIosVersion := queryValues.Get("wn_ios_version")

	// The URL of the API you want to call
	baseURL, err := url.Parse("https://test-fs.free.beeceptor.com/br")
	if err != nil {
		fmt.Println("Malformed URL:", err)
		return
	}

	// Prepare query parameters
	params := url.Values{}
	params.Add("branch_id", branchID)
	params.Add("requestID", requestID)
	params.Add("timezone", timezone)
	params.Add("client_detail", clientDetail)
	params.Add("wn_ios_version", wnIosVersion)

	// Encode the parameters and append to the URL
	baseURL.RawQuery = params.Encode()

	// Now baseURL.String() contains the full URL with the query parameters
	fmt.Println("Encoded URL:", baseURL.String())

	// Create a new POST request
	response, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println("Error fetching URL:", err)
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
	var responseBody dto.BrResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error Marshal": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responseBody)
}

func (p PServiceImpl) DO(c *gin.Context) {
	var requestBody dto.DRequest
	// Bind the JSON to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := json.Marshal(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error Marshal": err.Error()})
		return
	}

	data := url.Values{}
	data.Set("request_id", requestBody.RequestID)
	data.Set("owner_class_id", requestBody.OwnerClassID)
	data.Set("owner_id", requestBody.OwnerID)
	data.Set("restaurant_id", requestBody.RestaurantID)

	// The URL of the API you want to call
	apiURL := "https://test-fs.free.beeceptor.com/d"

	// Create a new POST request
	request, err := http.NewRequest("DELETE", apiURL, bytes.NewBufferString(data.Encode()))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to create request"})
		return
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

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
	var responseBody dto.DResponse
	err = json.Unmarshal(body, &responseBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error Marshal": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responseBody)
}

func (p PServiceImpl) SavePac(c *gin.Context) {
	var requestBody dto.SavePackageRequest
	// Bind the JSON to the struct
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// todo send to another API here

	responseBody := dto.SavePacResponse{
		Success:   "1",
		RequestID: requestBody.RequestID,
		PackageList: dto.Package{

			Name:        "Example Name",
			RequestID:   "REQ123456",
			PkgImgKey:   "ImgKey123",
			BranchID:    "Branch001",
			PackageID:   "Package123",
			Description: "Description of the package",
			OwnerID:     "Owner123",
			SubPkgMenuList: []dto.SubPkgMenu{
				{MenuID: 1, Name: "abc", Price: 100},
			},
			MenuList: []dto.Menu{
				{MenuID: 1, PkgMenuID: 1, DeleteFlag: false},
			},
		},
	}

	c.JSON(http.StatusOK, responseBody)
}
