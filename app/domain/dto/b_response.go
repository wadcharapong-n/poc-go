package dto

import "time"

type Employee struct {
	OwnerID        string    `json:"owner_id" example:"1"`
	Email          string    `json:"email" example:"john.doe@example.com"`
	ImageKey       string    `json:"image_key" example:"img123456"`
	FirstName      string    `json:"firstname" example:"John"`
	LastName       string    `json:"lastname" example:"Doe"`
	PhoneNumber    string    `json:"phonenumber" example:"555-1234"`
	UnderID        string    `json:"under_id" example:"2"`
	OwnerClassID   string    `json:"owner_class_id" example:"001"`
	OwnerClassName string    `json:"owner_class_name" example:"VIP"`
	StartTime      time.Time `json:"start_time" example:"2023-07-01T15:04:05Z07:00"`
	EndTime        time.Time `json:"end_time" example:"2023-07-03T15:04:05Z07:00"`
}

type Permission struct {
	ID   string `json:"id" example:"1"`
	Desc string `json:"desc" example:"permission 1"`
}

// BrResponse Define a struct for your response
type BrResponse struct {
	Success        string       `json:"success" example:"1"`
	RequestID      string       `json:"requestID" example:"REQ-001"`
	EmployeeList   []Employee   `json:"employee_list"`
	PermissionList []Permission `json:"permission_list"`
}
