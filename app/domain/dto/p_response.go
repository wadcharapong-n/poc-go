package dto

import "time"

// PaResponse Define a struct for your response
type PaResponse struct {
	Success         string    `json:"success" example:"1"`
	RequestID       string    `json:"requestID" example:"REQ-001"`
	PaymentID       string    `json:"payment_id" example:"PAY-123"`
	PmtTime         time.Time `json:"pmt_time" example:"2024-07-03T12:00:00Z"`
	ReceiptNo       string    `json:"receipt_no" example:"REC-456"`
	QrText          string    `json:"qr_text" example:"HEADER_QR-QR-789"`
	IsVoid          bool      `json:"is_void" example:"false"`
	ReviewPoint     int       `json:"review_point" example:"5"`
	OrderID         string    `json:"order_id" example:"ORDER-001"`
	PaymentRefID    string    `json:"payment_ref_id" example:"REF-002"`
	SplitPmtDetail  string    `json:"split_pmt_detail" example:"Split payment details here"`
	ResMemID        string    `json:"res_mem_id" example:"MEM-003"`
	MemberID        string    `json:"member_id" example:"MEM-004"`
	FlowAcctResp    string    `json:"flow_acct_resp" example:"Tax invoice response here"`
	FreshLivingResp string    `json:"freshLivingResp" example:"No Data"`
}
