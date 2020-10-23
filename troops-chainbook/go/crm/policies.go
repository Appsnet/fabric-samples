package main

// RiderProduct model
type RiderProduct struct {
	PolicyID    string
	ProductID   string
	PolicyNo    string
	ProductName string
}

// Policy model
type Policy struct {
	PolicyNo               string  `json:"policy_no"`
	OwnerChineseName       string  `json:"owner_chinese_name"`
	OwnerFirstName         string  `json:"owner_first_name"`
	OwnerLastName          string  `json:"owner_last_name"`
	OwnerDob               string  `json:"owner_dob"`
	OwnerEmail             string  `json:"owner_email"`
	InsuredChineseName     string  `json:"insured_chinese_name"`
	InsuredFirstName       string  `json:"insured_first_name"`
	InsuredLastName        string  `json:"insured_last_name"`
	InsuredDob             string  `json:"insured_dob"`
	InsuredEmail           string  `json:"insured_email"`
	ClientID               string  `json:"client_id"`
	InsuranceCompanyName   string  `json:"insurance_company_name"`
	ServiceName            string  `json:"service_name"`
	ServiceTel             string  `json:"service_tel"`
	CoopName               string  `json:"coop_name"`
	CoopTel                string  `json:"coop_tel"`
	SubmittedDate          string  `json:"submitted_date"`
	PolicyDate             string  `json:"policy_date"`
	PaymentType            string  `json:"payment_type"`
	CurrencyType           string  `json:"currency_type"`
	AppPreMiumAmount       float64 `json:"app_premium_amount"`
	PolicyAmount           float64 `json:"policy_amount"`
	SumAssuredAmount       float64 `json:"sum_assured_amount"`
	RecdRemiumAmount       float64 `json:"recd_remium_amount"`
	ProductID              string  `json:"product_id"`
	ProductName            string  `json:"product_name"`
	Period                 string  `json:"period"`
	CoolingOffDate         string  `json:"cooling_off_date"`
	FirstPremiumPaymentDay string  `json:"first_premium_payment_day"`
	ReceivePolicyDate      string  `json:"receive_policy_date"`
	PolicyStatus           int     `json:"policy_status"`
	CreateTime             string  `json:"create_time"`
	RiderProducts          string  `json:"rider_prodects"`
}

// Policies model
type Policies struct {
	Policies []Policy
}
