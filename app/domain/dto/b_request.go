package dto

type BRequest struct {
	BranchID     string `json:"branch_id" example:"123"`
	RequestID    string `json:"requestID" example:"456"`
	Timezone     string `json:"timezone" example:"UTC+0"`
	ClientDetail string `json:"client_detail" example:"Details about the client"`
	WnIOSVersion string `json:"wn_ios_version" example:"1.0.2"`
}
