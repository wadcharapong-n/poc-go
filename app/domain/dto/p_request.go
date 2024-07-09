package dto

import "time"

type Channel struct {
	OrderChannelID string `json:"order_channel_id"`
	ChannelOrderNo string `json:"channel_order_no"`
}

type OrderCustomerDetail struct {
	CustomerName string `json:"cust_name"`
	PhoneNumber  string `json:"phone_number"`
	Address      string `json:"address"`
	Comment      string `json:"comment"`
	PhonePrefix  string `json:"phone_prefix"`
	CountryCode  string `json:"country_code"`
}

type CommentObj struct {
	CompanyPrefix      string `json:"company_prefix"`
	CompanyCountryCode string `json:"company_country_code"`
	CompanyName        string `json:"company_name"`
	CompanyAddress     string `json:"company_address"`
	CompanyTaxID       string `json:"company_tax_id"`
	CompanyPhone       string `json:"company_phone"`
	CompanyBranch      string `json:"company_branch"`
	CompanyEmail       string `json:"company_email"`
}

type MposCreditDetail struct {
	StatusCode     string `json:"statusCode"`
	StatusMessage  string `json:"statusMessage"`
	ApprovalCode   string `json:"approvalCode"`
	TransactionID  string `json:"transactionID"`
	TraceNo        string `json:"traceNo"`
	CardNo         string `json:"cardNo"`
	CardType       string `json:"cardType"`
	CardHolderName string `json:"cardHolderName"`
	Reference      string `json:"reference"`
}

type OrderMenuList struct {
	IsCustomer        bool         `json:"is_customer"`
	ServedTime        time.Time    `json:"served_time"`
	ServedBy          string       `json:"served_by"`
	SendOrderSystemID string       `json:"send_order_system_id"`
	ActionBy          string       `json:"action_by"`
	OrderMenuID       string       `json:"order_menu_id"`
	MenuOrderTime     time.Time    `json:"menu_order_time"`
	MenuID            string       `json:"menu_id"`
	MenuGroupID       string       `json:"menu_group_id"`
	MenuGroupName     string       `json:"menu_group_name"`
	MenuName          string       `json:"menu_name"`
	MenuComment       string       `json:"menu_comment"`
	MenuPrice         float64      `json:"menu_price"`
	Quantity          int          `json:"quantity"`
	SendOrderBy       string       `json:"send_order_by"`
	OptionList        []OptionList `json:"option_list"`
	RejectTime        time.Time    `json:"reject_time"`
	RejectBy          string       `json:"reject_by"`
	RejectReason      string       `json:"reject_reason"`
	RejectFlag        bool         `json:"reject_flag"`
	CancelFlag        bool         `json:"cancel_flag"`
}

type OptionList struct {
	PaymentItemID string `json:"payment_item_id"`
	VoidFlag      bool   `json:"void_flag"`
	Quantity      int    `json:"quantity"`
	ActionBy      string `json:"action_by"`
	VoidReason    string `json:"void_reason"`
}

type PmtMenuList struct {
	ActionBy             string          `json:"action_by"`
	PaymentItemID        string          `json:"payment_item_id"`
	PromotionMenuID      string          `json:"promotion_menu_id"`
	PromotionMenuGroupID string          `json:"promotion_menu_group_id"`
	VoidFlag             bool            `json:"void_flag"`
	DiscountType         string          `json:"discount_type"`
	DiscountReason       string          `json:"discount_reason"`
	PromotionID          string          `json:"promotion_id"`
	OverridePromotion    bool            `json:"override_promotion"`
	ProUserID            string          `json:"pro_user_id"`
	MenuGroupID          MenuGroupID     `json:"menu_group_id"`
	MenuID               MenuID          `json:"menu_id"`
	IsNonVat             bool            `json:"is_non_vat"`
	PriceWithOption      float64         `json:"price_with_option"`
	PromotionRefCode     string          `json:"promotion_ref_code"`
	FsCrmPromotionUUID   string          `json:"fs_crm_promotion_uuid"`
	PromotionName        string          `json:"promotion_name"`
	MenuName             string          `json:"menu_name"`
	MenuGroupName        string          `json:"menu_group_name"`
	Quantity             int             `json:"quantity"`
	Price                float64         `json:"price"`
	DiscountedPrice      float64         `json:"discounted_price"`
	DiscountValue        float64         `json:"discount_value"`
	OrderMenuList        []OrderMenuList `json:"order_menu_list"`
	OptionList           []OptionList    `json:"option_list"`
}

type MenuGroupID struct {
	CatGrpID   string `json:"cat_grp_id"`
	CatGrpName string `json:"cat_grp_name"`
}

type MenuID struct {
	CostCenter string `json:"cost_center"`
	DeptCode   string `json:"dept_code"`
}

type QuestionList struct {
	QID        string       `json:"q_id"`
	ChoiceList []ChoiceList `json:"choice_list"`
}

type ChoiceList struct {
	CID   string `json:"c_id"`
	Count int    `json:"count"`
}

type SplitPaymentDetail struct {
	PaymentType     string  `json:"payment_type"`
	PaidAmount      float64 `json:"paid_amount"`
	SplitAmount     float64 `json:"split_amount"`
	ChangeAmount    float64 `json:"change_amount"`
	CreditCardID    string  `json:"credit_card_id"`
	PrefixCardDigit string  `json:"prefix_card_digit"`
	CardDigit       string  `json:"card_digit"`
	CustPmtRef      string  `json:"cust_pmt_ref"`
}

type PaRequest struct {
	Channel                 Channel              `json:"channel"`
	EdcSerialNo             string               `json:"edc_serial_no"`
	PaymentChannel          string               `json:"payment_channel"`
	BankCode                string               `json:"bank_code"`
	OrderType               string               `json:"order_type"`
	RequestID               string               `json:"requestID"`
	BranchID                string               `json:"branch_id"`
	TableName               string               `json:"table_name"`
	Timezone                string               `json:"timezone"`
	OrderCustomerDetail     OrderCustomerDetail  `json:"order_customer_detail"`
	OrderID                 string               `json:"order_id"`
	OpenOwnerID             string               `json:"open_owner_id"`
	CloseOwnerID            string               `json:"close_owner_id"`
	UniqueOrderCode         string               `json:"unique_order_code"`
	TableID                 string               `json:"table_id"`
	SeatAmt                 int                  `json:"seat_amt"`
	KioskType               string               `json:"kiosk_type"`
	OpenOrderDt             time.Time            `json:"open_order_dt"`
	ResvID                  string               `json:"resv_id"`
	CommentObj              CommentObj           `json:"comment_obj"`
	PaymentID               string               `json:"payment_id"`
	CancelBill              bool                 `json:"cancel_bill"`
	MposCreditDetail        MposCreditDetail     `json:"mpos_credit_detail"`
	BillDiscountPromotionID string               `json:"bill_discount_promotion_id"`
	ProUserID               string               `json:"pro_user_id"`
	CreditCardID            string               `json:"credit_card_id"`
	QrText                  string               `json:"qr_text"`
	CashDrawerCode          string               `json:"cash_drawer_code"`
	PrefixCardNo            string               `json:"prefix_cardNo"`
	CardNo                  string               `json:"cardNo"`
	BillDiscountType        string               `json:"bill_discount_type"`
	HostName                string               `json:"host_name"`
	ClientIP                string               `json:"client_ip"`
	ClientName              string               `json:"client_name"`
	RoundingAmount          float64              `json:"rounding_amount"`
	RoundingType            string               `json:"rounding_type"`
	BuffetPackageID         string               `json:"buffet_package_id"`
	CustPmtID               string               `json:"cust_pmt_id"`
	CustPmtRef              string               `json:"cust_pmt_ref"`
	RunningBook             string               `json:"running_book"`
	RunningNo               string               `json:"running_no"`
	PosID                   string               `json:"pos_id"`
	VoidAfterPrintInv       bool                 `json:"void_after_print_inv"`
	CancelMenuFlag          bool                 `json:"cancel_menu_flag"`
	Latitude                float64              `json:"latitude"`
	Longitude               float64              `json:"longitude"`
	TaxPercent              float64              `json:"tax_percent"`
	ServicePercent          float64              `json:"service_percent"`
	SubAmount               float64              `json:"sub_amount"`
	Service                 float64              `json:"service"`
	Tax                     float64              `json:"tax"`
	PmtApiRef               string               `json:"pmt_api_ref"`
	BeforeVat               float64              `json:"before_vat"`
	BeforeSvc               float64              `json:"before_svc"`
	NonVatAmount            float64              `json:"non_vat_amount"`
	PromotionRefCode        string               `json:"promotion_ref_code"`
	Version                 string               `json:"version"`
	CrmPhoneNo              string               `json:"crm_phone_no"`
	CrmName                 string               `json:"crm_name"`
	CrmRedeemID             string               `json:"crmRedeemId"`
	FsCrmHistoryID          string               `json:"fs_crm_history_id"`
	FsRedeemID              string               `json:"fs_redeem_id"`
	FsCrmMemberID           string               `json:"fs_crm_member_id"`
	CrmSystemID             string               `json:"crm_system_id"`
	MemberID                string               `json:"memberId"`
	FsCrmPromotionUUID      string               `json:"fs_crm_promotion_uuid"`
	PromotionName           string               `json:"promotion_name"`
	BillDiscountValue       float64              `json:"bill_discount_value"`
	BillDiscountedPrice     float64              `json:"bill_discounted_price"`
	BillDiscountReason      string               `json:"bill_discount_reason"`
	BillDiscountBy          string               `json:"bill_discount_by"`
	PaymentType             string               `json:"payment_type"`
	PaymentAmount           float64              `json:"payment_amount"`
	ChangeAmount            float64              `json:"change_amount"`
	RefundAmount            float64              `json:"refund_amount"`
	DiscountAmount          float64              `json:"discount_amount"`
	TotalAmount             float64              `json:"total_amount"`
	CreateDt                time.Time            `json:"create_dt"`
	ReceiptNo               string               `json:"receipt_no"`
	TipsAmount              float64              `json:"tips_amount"`
	CustomerName            string               `json:"customer_name"`
	PhoneNumber             string               `json:"phone_number"`
	PhonePrefix             string               `json:"phone_prefix"`
	CountryCode             string               `json:"country_code"`
	RestaurantID            string               `json:"restaurant_id"`
	SplitPaymentDetail      []SplitPaymentDetail `json:"split_payment_detail"`
	OrderMenuList           []OrderMenuList      `json:"order_menu_list"`
	PmtMenuList             []PmtMenuList        `json:"pmt_menu_list"`
	NoOfMenu                int                  `json:"no_of_menu"`
	QuestionList            []QuestionList       `json:"question_list"`
	TheOneCardCode          string               `json:"the_one_card_code"`
	BranchUUID              string               `json:"branch_uuid"`
}
