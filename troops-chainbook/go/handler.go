package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

// GetContract return policy contract form wallet
func GetContract() error {

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"troops-network",
		"organizations",
		"peerOrganizations",
		"org1.troops.io",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	Contract = network.GetContract("policy")
	return nil
}

// Contract of policy
var Contract *gateway.Contract

// GetAllCustomers from db
func GetAllCustomers(c *fiber.Ctx) error {

	return c.JSON(&fiber.Map{
		"success": true,
		// "result":  result,
		"message": "All customers returned",
	})
}

// GetCustomer from db
func GetCustomer(c *fiber.Ctx) error {
	email := c.Params("email")

	return c.JSON(&fiber.Map{
		"success": true,
		"result":  email,
		"message": "All customers returned",
	})
}

// CreatePolicy from db
func CreatePolicy(c *fiber.Ctx) error {

	result, err := Contract.SubmitTransaction("CreatePolicy",
		"POLICY_NO",
		"CLIENT_ID",
		"BROKER",
		"INSURANCE_COMPANY_NAME",
		"COOP_NAME",
		"SUBMITTED_DATE",
		"POLICY_DATE",
		"PAYMENT_TYPE",
		"CURRENCY_TYPE",
		"0",
		"0",
		"0",
		"PRODUCT_ID",
		"PRODUCT_NAME",
		"PERIOD",
		"COOLING_OFF_DATE",
		"FIRST_PREMIUM_PAYMENT_DAY",
		"RECEIVE_POLICY_DATE",
		"CREATE_TIME",
	)
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	return c.JSON(&fiber.Map{
		"success": true,
		"result":  result,
		"message": "all policies returned",
	})
}

// GetAllPolicies from db
func GetAllPolicies(c *fiber.Ctx) error {
	result, err := Contract.EvaluateTransaction("GetAllPolicy")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))
	var policies []Policy
	json.Unmarshal([]byte(result), &policies)

	return c.JSON(&fiber.Map{
		"success": true,
		"result":  policies,
		"message": "all policies returned",
	})
}

// GetPolicy query policy by policy no
func GetPolicy(c *fiber.Ctx) error {
	// policyno := c.Params("policyno")
	result, err := Contract.EvaluateTransaction("ReadPolicy", "policyno")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	return c.JSON(&fiber.Map{
		"success": true,
		"result":  result,
		"message": "Policy returned",
	})
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"..",
		"troops-network",
		"organizations",
		"peerOrganizations",
		"org1.troops.io",
		"users",
		"User1@org1.troops.io",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "User1@org1.troops.io-cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
}
