package dto

import "time"

type Channel struct {
	OrderChannelID string `json:"order_channel_id" example:"123"`
	ChannelOrderNo string `json:"channel_order_no" example:"456"`
}

type OrderMenu struct {
	IsCustomer        bool      `json:"is_customer" example:"true"`
	ServedTime        time.Time `json:"served_time" example:"2024-07-03T12:15:00Z"`
	ServedBy          string    `json:"served_by" example:"Server1"`
	SendOrderSystemID string    `json:"send_order_system_id" example:"SYS-001"`
	ActionBy          string    `json:"action_by" example:"Chef1"`
	OrderMenuID       string    `json:"order_menu_id" example:"OM-001"`
	MenuOrderTime     time.Time `json:"menu_order_time" example:"2024-07-03T12:05:00Z"`
	MenuID            string    `json:"menu_id" example:"MENU-001"`
	MenuGroupID       string    `json:"menu_group_id" example:"MG-001"`
	MenuGroupName     string    `json:"menu_group_name" example:"Appetizers"`
	MenuName          string    `json:"menu_name" example:"Spring Rolls"`
	MenuComment       string    `json:"menu_comment" example:"Extra crispy"`
	MenuPrice         float64   `json:"menu_price" example:"5.0"`
	Quantity          int       `json:"quantity" example:"2"`
	SendOrderBy       string    `json:"send_order_by" example:"Chef1"`
	OptionList        []Option  `json:"option_list"`
	RejectTime        time.Time `json:"reject_time" example:"2024-07-03T12:10:00Z"`
	RejectBy          string    `json:"reject_by" example:"Chef2"`
	RejectReason      string    `json:"reject_reason" example:"Out of stock"`
	RejectFlag        bool      `json:"reject_flag" example:"false"`
	CancelFlag        bool      `json:"cancel_flag" example:"false"`
}

type Option struct {
	PaymentItemID string `json:"payment_item_id" example:"PI-001"`
	VoidFlag      bool   `json:"void_flag" example:"false"`
	Quantity      int    `json:"quantity" example:"1"`
	ActionBy      string `json:"action_by" example:"Chef1"`
	VoidReason    string `json:"void_reason" example:"None"`
}

type PmtMenu struct {
	ActionBy             string      `json:"action_by" example:"Admin"`
	PaymentItemID        string      `json:"payment_item_id" example:"PMT-001"`
	PromotionMenuID      string      `json:"promotion_menu_id" example:"PROMO-001"`
	PromotionMenuGroupID string      `json:"promotion_menu_group_id" example:"PMG-001"`
	VoidFlag             bool        `json:"void_flag" example:"false"`
	DiscountType         string      `json:"discount_type" example:"Percentage"`
	DiscountReason       string      `json:"discount_reason" example:"Loyalty Discount"`
	PromotionID          string      `json:"promotion_id" example:"PROMO-002"`
	OverridePromotion    bool        `json:"override_promotion" example:"true"`
	ProUserID            string      `json:"pro_user_id" example:"PRO-002"`
	MenuGroupID          MenuGroupID `json:"menu_group_id"`
	MenuID               MenuID      `json:"menu_id"`
	IsNonVat             bool        `json:"is_non_vat" example:"false"`
	PriceWithOption      float64     `json:"price_with_option" example:"10.0"`
	PromotionRefCode     string      `json:"promotion_ref_code" example:"PROMO-REF-001"`
	FsCrmPromotionUuid   string      `json:"fs_crm_promotion_uuid" example:"FS-CRM-PROMO-UUID-002"`
	PromotionName        string      `json:"promotion_name" example:"Holiday Discount"`
	MenuName             string      `json:"menu_name" example:"Pizza"`
	MenuGroupName        string      `json:"menu_group_name" example:"Main Course"`
	Quantity             int         `json:"quantity" example:"1"`
	Price                float64     `json:"price" example:"10.0"`
	DiscountedPrice      float64     `json:"discounted_price" example:"9.0"`
	DiscountValue        float64     `json:"discount_value" example:"1.0"`
	VoidReason           string      `json:"void_reason" example:"None"`
	OrderMenuList        []OrderMenu `json:"order_menu_list"`
	OptionList           []Option    `json:"option_list"`
}

type MenuGroupID struct {
	CatGrpID   string `json:"cat_grp_id" example:"CAT-001"`
	CatGrpName string `json:"cat_grp_name" example:"Category1"`
}

type MenuID struct {
	CostCenter string `json:"cost_center" example:"CC-001"`
	DeptCode   string `json:"dept_code" example:"DPT-001"`
}

type Question struct {
	QID        string   `json:"q_id" example:"Q-001"`
	ChoiceList []Choice `json:"choice_list"`
}

type Choice struct {
	CID   string `json:"c_id" example:"C-001"`
	Count int    `json:"count" example:"5"`
}

type OrderCustomerDetail struct {
	CustName    string `json:"cust_name" example:"John Doe"`
	PhoneNumber string `json:"phone_number" example:"1234567890"`
	Address     string `json:"address" example:"123 Street Name, City, Country"`
	Comment     string `json:"comment" example:"No onions please"`
	PhonePrefix string `json:"phone_prefix" example:"+1"`
	CountryCode string `json:"country_code" example:"US"`
}

type CommentObj struct {
	CompanyPrefix      string `json:"company_prefix" example:"COMP"`
	CompanyCountryCode string `json:"company_country_code" example:"US"`
	CompanyName        string `json:"company_name" example:"Company Inc."`
	CompanyAddress     string `json:"company_address" example:"456 Business Ave, City, Country"`
	CompanyTaxID       string `json:"company_tax_id" example:"TAX123"`
	CompanyPhone       string `json:"company_phone" example:"0987654321"`
	CompanyBranch      string `json:"company_branch" example:"Branch 1"`
	CompanyEmail       string `json:"company_email" example:"info@company.com"`
}

type MposCreditDetail struct {
	StatusCode     string `json:"statusCode" example:"200"`
	StatusMessage  string `json:"statusMessage" example:"Approved"`
	ApprovalCode   string `json:"approvalCode" example:"APPROVED"`
	TransactionID  string `json:"transactionID" example:"TXN-001"`
	TraceNo        string `json:"traceNo" example:"TRACE-001"`
	CardNo         string `json:"cardNo" example:"4111111111111111"`
	CardType       string `json:"cardType" example:"VISA"`
	CardHolderName string `json:"cardHolderName" example:"John Doe"`
	Reference      string `json:"reference" example:"REF-001"`
}

type SplitPaymentDetail struct {
	PaymentType     string  `json:"payment_type" example:"Cash"`
	PaidAmount      float64 `json:"paid_amount" example:"50.0"`
	SplitAmount     float64 `json:"split_amount" example:"40.0"`
	ChangeAmount    float64 `json:"change_amount" example:"10.0"`
	CreditCardID    string  `json:"credit_card_id" example:"CC-002"`
	PrefixCardDigit string  `json:"prefix_card_digit" example:"1234"`
	CardDigit       string  `json:"card_digit" example:"5678"`
	CustPmtRef      string  `json:"cust_pmt_ref" example:"CPMT-REF-002"`
}

// PaRequest represents the JSON structure for a payment request.
type PaRequest struct {
	Channel                 Channel              `json:"channel"`
	EdcSerialNo             string               `json:"edc_serial_no" example:"789"`
	PaymentChannel          string               `json:"payment_channel" example:"Credit Card"`
	BankCode                string               `json:"bank_code" example:"XYZ"`
	OrderType               string               `json:"order_type" example:"Dine-In"`
	RequestID               string               `json:"requestID" example:"REQ-001"`
	BranchID                string               `json:"branch_id" example:"BR-001"`
	TableName               string               `json:"table_name" example:"Table 5"`
	Timezone                string               `json:"timezone" example:"GMT+7"`
	OrderCustomerDetail     OrderCustomerDetail  `json:"order_customer_detail"`
	OrderID                 string               `json:"order_id" example:"ORDER-001"`
	OpenOwnerID             string               `json:"open_owner_id" example:"OWNER-001"`
	CloseOwnerID            string               `json:"close_owner_id" example:"OWNER-002"`
	UniqueOrderCode         string               `json:"unique_order_code" example:"UNIQUE-123"`
	TableID                 string               `json:"table_id" example:"TBL-001"`
	SeatAmt                 int                  `json:"seat_amt" example:"4"`
	KioskType               string               `json:"kiosk_type" example:"Self-Service"`
	OpenOrderDT             time.Time            `json:"open_order_dt" example:"2024-07-03T12:00:00Z"`
	ResvID                  string               `json:"resv_id" example:"RESV-001"`
	CommentObj              CommentObj           `json:"comment_obj"`
	PaymentID               string               `json:"payment_id" example:"PMT-001"`
	CancelBill              bool                 `json:"cancel_bill" example:"false"`
	MposCreditDetail        MposCreditDetail     `json:"mpos_credit_detail"`
	BillDiscountPromotionID string               `json:"bill_discount_promotion_id" example:"BDP-001"`
	ProUserID               string               `json:"pro_user_id" example:"PRO-001"`
	CreditCardID            string               `json:"credit_card_id" example:"CC-001"`
	QrText                  string               `json:"qr_text" example:"QR-001"`
	CashDrawerCode          string               `json:"cash_drawer_code" example:"CD-001"`
	PrefixCardNo            string               `json:"prefix_cardNo" example:"4111"`
	CardNo                  string               `json:"cardNo" example:"1111"`
	BillDiscountType        string               `json:"bill_discount_type" example:"Percentage"`
	HostName                string               `json:"host_name" example:"localhost"`
	ClientIP                string               `json:"client_ip" example:"192.168.1.1"`
	ClientName              string               `json:"client_name" example:"Client1"`
	RoundingAmount          float64              `json:"rounding_amount" example:"0.05"`
	RoundingType            string               `json:"rounding_type" example:"Up"`
	BuffetPackageID         string               `json:"buffet_package_id" example:"BP-001"`
	CustPmtID               string               `json:"cust_pmt_id" example:"CPMT-001"`
	CustPmtRef              string               `json:"cust_pmt_ref" example:"CPMT-REF-001"`
	RunningBook             string               `json:"running_book" example:"RB-001"`
	RunningNo               string               `json:"running_no" example:"RN-001"`
	PosID                   string               `json:"pos_id" example:"POS-001"`
	VoidAfterPrintInv       bool                 `json:"void_after_print_inv" example:"false"`
	CancelMenuFlag          bool                 `json:"cancel_menu_flag" example:"false"`
	Latitude                float64              `json:"latitude" example:"40.7128"`
	Longitude               float64              `json:"longitude" example:"-74.0060"`
	TaxPercent              float64              `json:"tax_percent" example:"7.0"`
	ServicePercent          float64              `json:"service_percent" example:"10.0"`
	SubAmount               float64              `json:"sub_amount" example:"100.0"`
	Service                 float64              `json:"service" example:"10.0"`
	Tax                     float64              `json:"tax" example:"7.0"`
	PmtApiRef               string               `json:"pmt_api_ref" example:"PMT-API-001"`
	BeforeVat               float64              `json:"before_vat" example:"93.0"`
	BeforeSvc               float64              `json:"before_svc" example:"100.0"`
	NonVatAmount            float64              `json:"non_vat_amount" example:"0.0"`
	PromotionRefCode        string               `json:"promotion_ref_code" example:"PROMO-001"`
	Version                 string               `json:"version" example:"1.0"`
	CrmPhoneNo              string               `json:"crm_phone_no" example:"1234567890"`
	CrmName                 string               `json:"crm_name" example:"John Doe"`
	CrmRedeemId             string               `json:"crmRedeemId" example:"CRM-REDEEM-001"`
	FsCrmHistoryId          string               `json:"fs_crm_history_id" example:"FS-CRM-HIST-001"`
	FsRedeemId              string               `json:"fs_redeem_id" example:"FS-REDEEM-001"`
	FsCrmMemberId           string               `json:"fs_crm_member_id" example:"FS-CRM-MEM-001"`
	CrmSystemId             string               `json:"crm_system_id" example:"CRM-SYS-001"`
	MemberId                string               `json:"memberId" example:"MEM-001"`
	FsCrmPromotionUuid      string               `json:"fs_crm_promotion_uuid" example:"FS-CRM-PROMO-UUID-001"`
	PromotionName           string               `json:"promotion_name" example:"Summer Sale"`
	BillDiscountValue       float64              `json:"bill_discount_value" example:"10.0"`
	BillDiscountedPrice     float64              `json:"bill_discounted_price" example:"90.0"`
	BillDiscountReason      string               `json:"bill_discount_reason" example:"Seasonal Discount"`
	BillDiscountBy          string               `json:"bill_discount_by" example:"Admin"`
	PaymentType             string               `json:"payment_type" example:"Credit Card"`
	PaymentAmount           float64              `json:"payment_amount" example:"90.0"`
	ChangeAmount            float64              `json:"change_amount" example:"0.0"`
	RefundAmount            float64              `json:"refund_amount" example:"0.0"`
	DiscountAmount          float64              `json:"discount_amount" example:"10.0"`
	TotalAmount             float64              `json:"total_amount" example:"100.0"`
	CreateDt                time.Time            `json:"create_dt" example:"2024-07-03T12:00:00Z"`
	ReceiptNo               string               `json:"receipt_no" example:"REC-001"`
	TipsAmount              float64              `json:"tips_amount" example:"5.0"`
	CustomerName            string               `json:"customer_name" example:"John Doe"`
	PhoneNumber             string               `json:"phone_number" example:"1234567890"`
	PhonePrefix             string               `json:"phone_prefix" example:"+1"`
	CountryCode             string               `json:"country_code" example:"US"`
	RestaurantID            string               `json:"restaurant_id" example:"RES-001"`
	SplitPaymentDetail      []SplitPaymentDetail `json:"split_payment_detail"`
	OrderMenuList           []OrderMenu          `json:"order_menu_list"`
	PmtMenuList             []PmtMenu            `json:"pmt_menu_list"`
	NoOfMenu                int                  `json:"no_of_menu" example:"2"`
	QuestionList            []Question           `json:"question_list"`
	TheOneCardCode          string               `json:"the_one_card_code" example:"ONE-001"`
	BranchUuid              string               `json:"branch_uuid" example:"BR-UUID-001"`
}
