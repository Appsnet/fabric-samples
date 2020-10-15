package main

// RiderProduct model
type RiderProduct struct {
	PolicyID    string
	ProductID   string
	PolicyNo    string
	ProductName string
}

// // Policy model
// type Policy struct {
// 	id                     int
// 	PolicyNo               string
// 	OwnerChineseName       string
// 	OwnerFirstName         string
// 	OwnerLastName          string
// 	OwnerDob               string
// 	OwnerEmail             string
// 	InsuredChineseName     string
// 	InsuredFirstName       string
// 	InsuredLastName        string
// 	InsuredDob             string
// 	InsuredEmail           string
// 	ClientID               int
// 	InsuranceCompanyName   string
// 	ServiceName            string
// 	ServiceTel             string
// 	CoopName               string
// 	CoopTel                string
// 	SubmittedDate          string
// 	PolicyDate             string
// 	PaymentType            int
// 	CurrencyType           string
// 	AppPreMiumAmount       float32
// 	PolicyAmount           float32
// 	SumAssuredAmount       float32
// 	RecdRemiumAmount       float32
// 	ProductID              int
// 	ProductName            string
// 	Period                 string
// 	CoolingOffDate         string
// 	FirstPremiumPaymentDay string
// 	ReceivePolicyDate      string
// 	PolicyStatus           int
// 	CreateTime             string
// 	RiderProducts          []RiderProduct
// }

// Policy model
type Policy struct {
	PolicyNo               string `json:"policy_no"`
	ClientID               string `json:"client_id"`
	Broker                 string `json:"broker"`
	InsuranceCompanyName   string `json:"insurer"`
	CoopName               string `json:"coop"`
	PolicyDate             string `json:"policy_date"`
	PaymentType            string `json:"payment_type"`
	CurrencyType           string `json:"currency"`
	AppPremiumAmount       int    `json:"approved_policy_amount"`
	PolicyAmount           int    `json:"policy_amount"`
	RecdRemiumAmount       int    `json:"received_amount"`
	ProductID              string `json:"product_id"`
	ProductName            string `json:"product_name"`
	Period                 string `json:"peroid"`
	CoolingOffDate         string `json:"cooling_off_date"`
	FirstPremiumPaymentDay string `json:"first_payment_date"`
	ReceivePolicyDate      string `json:"receive_policy_date"`
	CreateTime             string `json:"create_timestamp"`
}

// Policies model
type Policies struct {
	Policies []Policy
}
