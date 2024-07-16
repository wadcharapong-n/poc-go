package dto

// DRequest RequestData represents the structure for request parameters.
type DRequest struct {
	OwnerClassID string `json:"owner_class_id" example:"123"`
	RequestID    string `json:"request_id" example:"456"`
	OwnerID      string `json:"owner_id" example:"789"`
	RestaurantID string `json:"restaurant_id" example:"101"`
}
