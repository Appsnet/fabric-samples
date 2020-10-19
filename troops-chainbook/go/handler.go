package main

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"

	"chainbook/fabric"
	"chainbook/models"
)

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

	result, err := fabric.Contract.SubmitTransaction("CreatePolicy",
		"00001",
		"OwnerChineseName",
		"OwnerFirstName",
		"OwnerLastName",
		"OwnerDob",
		"OwnerEmail",
		"InsuredChineseName",
		"InsuredFirstName",
		"InsuredLastName",
		"InsuredDob",
		"InsuredEmail",
		"ClientID",
		"InsuranceCompanyName",
		"ServiceName",
		"ServiceTel",
		"CoopName",
		"CoopTel",
		"SubmittedDate",
		"PolicyDate",
		"PaymentType",
		"CurrencyType",
		"0",
		"0",
		"0",
		"0",
		"ProductID",
		"ProductName",
		"Period",
		"CoolingOffDate",
		"FirstPremiumPaymentDay",
		"ReceivePolicyDate",
		"0",
		"CreateTime",
		"{}",
	)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Failed to evaluate transaction: %s\n", err),
		})
	}

	return c.JSON(&fiber.Map{
		"success": true,
		"result":  result,
		"message": "all policies returned",
	})
}

// GetAllPolicies from db
func GetAllPolicies(c *fiber.Ctx) error {
	var policies []models.Policy

	result, err := fabric.Contract.EvaluateTransaction("GetAllPolicy")
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Failed to evaluate transaction: %s\n", err),
		})
	}

	json.Unmarshal([]byte(result), &policies)

	return c.JSON(&fiber.Map{
		"success":       true,
		"result":        policies,
		"nubers_policy": len(policies),
		"message":       "all policies returned",
	})
}

// GetPolicy query policy by policy no
func GetPolicy(c *fiber.Ctx) error {
	policyno := c.Params("policyno")

	var policy models.Policy

	result, err := fabric.Contract.EvaluateTransaction("ReadPolicy", policyno)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   fmt.Sprintf("Failed to evaluate transaction: %s\n", err),
		})

	}

	json.Unmarshal([]byte(result), &policy)

	return c.JSON(&fiber.Map{
		"success": true,
		"result":  policy,
		"message": "Policy returned",
	})
}
