package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"

	"chainbook/conf"
)

// ImportPublicPolicy from db
func ImportPublicPolicy(contract *gateway.Contract, policy Policy) {

	result, err := contract.SubmitTransaction("CreatePolicy",
		policy.PolicyNo,
		policy.OwnerDob,
		policy.InsuredDob,
		policy.InsuranceCompanyName,
		policy.ServiceName,
		policy.ServiceTel,
		policy.CoopName,
		policy.CoopTel,
		policy.SubmittedDate,
		policy.PolicyDate,
		policy.PaymentType,
		policy.CurrencyType,
		fmt.Sprintf("%f", policy.AppPreMiumAmount),
		fmt.Sprintf("%f", policy.PolicyAmount),
		fmt.Sprintf("%f", policy.SumAssuredAmount),
		fmt.Sprintf("%f", policy.RecdRemiumAmount),
		policy.ProductID,
		policy.ProductName,
		policy.Period,
		policy.CoolingOffDate,
		policy.FirstPremiumPaymentDay,
		policy.ReceivePolicyDate,
		fmt.Sprintf("%d", policy.PolicyStatus),
		policy.CreateTime,
		policy.RiderProducts,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}

// ImportCustomer from db
func ImportCustomer(contract *gateway.Contract, policy Policy) {

	result, err := contract.SubmitTransaction("CreatePolicy",
		policy.PolicyNo,
		policy.OwnerChineseName,
		policy.OwnerFirstName,
		policy.OwnerLastName,
		policy.OwnerDob,
		policy.OwnerEmail,
		policy.InsuredChineseName,
		policy.InsuredFirstName,
		policy.InsuredLastName,
		policy.InsuredDob,
		policy.InsuredEmail,
		policy.ClientID,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}

// ImportPolicy from db
func ImportPolicy(contract *gateway.Contract, policy Policy) {

	result, err := contract.SubmitTransaction("CreatePolicy",
		policy.PolicyNo,
		policy.OwnerChineseName,
		policy.OwnerFirstName,
		policy.OwnerLastName,
		policy.OwnerDob,
		policy.OwnerEmail,
		policy.InsuredChineseName,
		policy.InsuredFirstName,
		policy.InsuredLastName,
		policy.InsuredDob,
		policy.InsuredEmail,
		policy.ClientID,
		policy.InsuranceCompanyName,
		policy.ServiceName,
		policy.ServiceTel,
		policy.CoopName,
		policy.CoopTel,
		policy.SubmittedDate,
		policy.PolicyDate,
		policy.PaymentType,
		policy.CurrencyType,
		fmt.Sprintf("%f", policy.AppPreMiumAmount),
		fmt.Sprintf("%f", policy.PolicyAmount),
		fmt.Sprintf("%f", policy.SumAssuredAmount),
		fmt.Sprintf("%f", policy.RecdRemiumAmount),
		policy.ProductID,
		policy.ProductName,
		policy.Period,
		policy.CoolingOffDate,
		policy.FirstPremiumPaymentDay,
		policy.ReceivePolicyDate,
		fmt.Sprintf("%d", policy.PolicyStatus),
		policy.CreateTime,
		policy.RiderProducts,
	)
	if err != nil {
		log.Println(err)
	}
	log.Println(result)
}

//GetPolicyByPolicyNo query policy by policy number
func GetPolicyByPolicyNo() error {

	mychannelContract, err := GetContract("private-channel", "policy-private")
	if err != nil {
		log.Println(err)
	}
	publicContract, err := GetContract("public-channel", "policy-public")
	if err != nil {
		log.Println(err)
	}
	brokerContract, err := GetContract("broker-channel", "policy-broker")
	if err != nil {
		log.Println(err)
	}

	rows, err := conf.DB.Query(fmt.Sprintf(`SELECT
    IFNULL(T_POLICY.POLICY_NO, ''),
    IFNULL(OWNER.CHINESE_NAME, '') AS OWNER_CHINESE_NAME,
    IFNULL(OWNER.FIRST_NAME, '') AS OWNER_FIRST_NAME,
    IFNULL(OWNER.LAST_NAME, '') AS OWNER_LAST_NAME,
    IFNULL(OWNER.DOB, '') AS OWNER_DOB,
    IFNULL(OWNER.EMAIL, '') AS OWNER_EMAIL,
    IFNULL(INSURED.CHINESE_NAME, '') AS INSURED_CHINESE_NAME,
    IFNULL(INSURED.FIRST_NAME, '') AS INSURED_FIRST_NAME,
    IFNULL(INSURED.LAST_NAME, '') AS INSURED_LAST_NAME,
    IFNULL(INSURED.DOB, '') AS INSURED_DOB,
    IFNULL(INSURED.EMAIL, '') AS INSURED_EMAIL,
    IFNULL(T_POLICY_CLIENT.CLIENT_ID, 0),
    IFNULL(T_INSURANCE_COMPANY.ENG_NAME, '') AS INSURANCE_COMPANY_NAME,
    IFNULL(SERVICE.NAME, '') AS SERVICE_NAME,
    IFNULL(SERVICE.CONTACT_NUMBER_COMPANY, '') AS SERVICE_TEL,
    IFNULL(T_STAFF.NAME, '') AS COOP_NAME,
    IFNULL(T_STAFF.CONTACT_NUMBER_COMPANY, '') AS COOP_TEL,
    IFNULL(T_POLICY.SUBMITTED_DATE, ''),
    IFNULL(T_POLICY.POLICY_DATE, ''),
    IFNULL(T_POLICY.PAYMENT_MODE, 0) AS PAYMENT_TYPE,
    IFNULL(T_CURRENCY.CURRENCY_TYPE, ''),
    IFNULL(T_POLICY.APP_PREMIUM_AMOUNT, 0),
	IFNULL(T_POLICY.APP_PREMIUM_AMOUNT, 0) AS POLICY_AMOUNT,
	IFNULL(T_POLICY.SUM_ASSURED_AMOUNT, 0),
    IFNULL(T_POLICY.RECD_REMIUM_AMOUNT, 0),
    IFNULL(T_PRODUCT.ID, 0) AS PRODUCT_ID,
    IFNULL(T_PRODUCT.NAME, '') AS PRODUCT_NAME,
    IFNULL(T_PRODUCT_PERIOD.NAME, '') AS PERIOD,
    IFNULL(T_POLICY.COOLING_OFF_DATE, ''),
    IFNULL(T_POLICY.FIRST_PREMIUM_PAYMENT_DAY, ''),
	IFNULL(T_POLICY.RECEIVE_POLICY_DATE, ''),
    IFNULL(T_POLICY.POLICY_STATUS, 0),
    IFNULL(T_POLICY.CREATE_TIME, '')
FROM
    T_POLICY_CLIENT
        LEFT JOIN
    T_POLICY ON T_POLICY_CLIENT.POLICY_ID = T_POLICY.ID
        LEFT JOIN
	T_PRODUCT ON T_POLICY.PLAN_ID = T_PRODUCT.ID
        LEFT JOIN
    T_INSURANCE_COMPANY ON T_INSURANCE_COMPANY.ID = T_PRODUCT.INSURANCE_COMPANY_ID
        LEFT JOIN
    T_STAFF ON T_POLICY.AGENT_ID = T_STAFF.ID
        LEFT JOIN
    T_STAFF AS SERVICE ON T_POLICY.AGENT_ID2 = SERVICE.ID
        LEFT JOIN
    T_CURRENCY ON T_POLICY.APP_PREMIUM_CURRENCY = T_CURRENCY.ID
        LEFT JOIN
    T_PRODUCT_PERIOD ON T_POLICY.PERIOD_ID = T_PRODUCT_PERIOD.ID
        LEFT JOIN
    T_POLICY_RIDER ON T_POLICY.ID = T_POLICY_RIDER.POLICY_ID
        LEFT JOIN
    (SELECT 
        POLICY_ID, CLIENT_ID, OWNER, INSURED
    FROM
        T_POLICY_CLIENT
    WHERE
        OWNER = 1) AS POLICY_OWNER ON POLICY_OWNER.POLICY_ID = T_POLICY.ID
        LEFT JOIN
    T_CUSTOMER AS OWNER ON POLICY_OWNER.CLIENT_ID = OWNER.ID
        LEFT JOIN
    (SELECT 
        POLICY_ID, CLIENT_ID, OWNER, INSURED
    FROM
        T_POLICY_CLIENT
    WHERE
        INSURED = 1) AS POLICY_INSURED ON POLICY_INSURED.POLICY_ID = T_POLICY.ID
        LEFT JOIN
    T_CUSTOMER AS INSURED ON POLICY_INSURED.CLIENT_ID = INSURED.ID
GROUP BY T_POLICY.POLICY_NO;`))
	if err != nil {
		log.Println(err.Error())
	}

	defer rows.Close()
	i := 1
	for rows.Next() {

		var policy Policy

		err = rows.Scan(
			&policy.PolicyNo,
			&policy.OwnerChineseName,
			&policy.OwnerFirstName,
			&policy.OwnerLastName,
			&policy.OwnerDob,
			&policy.OwnerEmail,
			&policy.InsuredChineseName,
			&policy.InsuredFirstName,
			&policy.InsuredLastName,
			&policy.InsuredDob,
			&policy.InsuredEmail,
			&policy.ClientID,
			&policy.InsuranceCompanyName,
			&policy.ServiceName,
			&policy.ServiceTel,
			&policy.CoopName,
			&policy.CoopTel,
			&policy.SubmittedDate,
			&policy.PolicyDate,
			&policy.PaymentType,
			&policy.CurrencyType,
			&policy.AppPreMiumAmount,
			&policy.PolicyAmount,
			&policy.SumAssuredAmount,
			&policy.RecdRemiumAmount,
			&policy.ProductID,
			&policy.ProductName,
			&policy.Period,
			&policy.CoolingOffDate,
			&policy.FirstPremiumPaymentDay,
			&policy.ReceivePolicyDate,
			&policy.PolicyStatus,
			&policy.CreateTime)
		if err != nil {
			log.Println(err.Error())
		}

		productsRows, err := conf.DB.Query(fmt.Sprintf(`
		SELECT 
			POLICY_ID,
			PRODUCT_ID,
			T_POLICY.POLICY_NO,
			T_PRODUCT.NAME AS PRODUCT_NAME
		FROM
			T_POLICY_RIDER
				LEFT JOIN
			T_POLICY ON T_POLICY_RIDER.POLICY_ID = T_POLICY.ID
				LEFT JOIN
			T_PRODUCT ON T_POLICY_RIDER.PRODUCT_ID = T_PRODUCT.ID
		WHERE
			T_POLICY.POLICY_NO = '%s';
		`, policy.PolicyNo))
		if err != nil {
			log.Println(err.Error())
		}

		var riderProducts []RiderProduct
		defer productsRows.Close()
		for productsRows.Next() {
			var product RiderProduct

			err = productsRows.Scan(&product.PolicyID, &product.ProductID, &product.PolicyNo, &product.ProductName)
			if err != nil {
				log.Println(err.Error())
			}

			// log.Println(product.ProductName)
			riderProducts = append(riderProducts, product)
		}
		if len(riderProducts) > 0 {
			riders, err := json.Marshal(riderProducts)
			if err != nil {
				log.Println(err)
			}
			policy.RiderProducts = string(riders)
		} else {
			policy.RiderProducts = "[]"
		}

		if policy.PolicyNo != "" {
			log.Println(i, policy.PolicyNo, policy.OwnerChineseName, policy.OwnerFirstName, policy.OwnerLastName, policy.CurrencyType)
			i++
			ImportPolicy(mychannelContract, policy)
			ImportCustomer(brokerContract, policy)
			ImportPublicPolicy(publicContract, policy)
		}
	}
	return nil
}

func main() {
	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	// Connect to database
	if err := conf.Connect(); err != nil {
		log.Fatal(err)
	}

	GetPolicyByPolicyNo()
}
