package dto

type PaResponse struct {
	Success         string `json:"success"`
	RequestID       string `json:"requestID"`
	PaymentID       string `json:"payment_id"`
	PmtTime         string `json:"pmt_time"`
	ReceiptNo       string `json:"receipt_no"`
	QRText          string `json:"qr_text"`
	IsVoid          bool   `json:"is_void"`
	ReviewPoint     int    `json:"review_point"`
	OrderID         string `json:"order_id"`
	PaymentRefID    string `json:"payment_ref_id"`
	SplitPmtDetail  string `json:"split_pmt_detail"`
	ResMemID        string `json:"res_mem_id"`
	MemberID        string `json:"member_id"`
	FlowAcctResp    string `json:"flow_acct_resp"`
	FreshLivingResp string `json:"freshLivingResp"`
}
