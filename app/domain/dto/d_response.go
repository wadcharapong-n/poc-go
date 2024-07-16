package dto

// DResponse Define a struct for your response
type DResponse struct {
	Success   string `json:"success" example:"1"`
	RequestID string `json:"requestID" example:"REQ-001"`
	Message   string `json:"message" example:"Deleted"`
}
