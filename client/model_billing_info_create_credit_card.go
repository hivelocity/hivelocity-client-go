/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BillingInfoCreateCreditCard struct {
	// The first name of the individual associated with the billing ID
	First string `json:"first"`
	// The street address associated with the billing ID
	Address string `json:"address"`
	// Make this the default payment method for future orders
	SetDefault bool `json:"setDefault,omitempty"`
	// Make this the default payment method for all services in the client's account
	DefaultAllServices bool `json:"defaultAllServices,omitempty"`
	// The state associated with the billing ID
	State string `json:"state"`
	// The zip/postal code associated with the billing ID
	Zip string `json:"zip"`
	// The last name of the individual associated with the billing ID
	Last string `json:"last"`
	// The company name associated with the billing ID
	Company string `json:"company,omitempty"`
	// The email address associated with the billing ID
	Email string `json:"email"`
	// The credit card network associated with the billing ID: `visa`, `mastercard`, `discover`, `amex`, etc.
	CcType string `json:"ccType"`
	// The city associated with the billing ID in question
	City string `json:"city"`
	// The expiration month for the credit card
	CcExpireMonth string `json:"ccExpireMonth"`
	// The expiration year for the credit card with 2 digits
	CcExpireYear string `json:"ccExpireYear"`
	// The identification number for the client associated with the Billing Info ID
	ClientId int32 `json:"clientId,omitempty"`
	// The credit card account number
	CcNum string `json:"ccNum"`
	// The credit card verification number
	CcCvv2 string `json:"ccCvv2"`
	// The phone number associated with the billing ID
	Phone string `json:"phone"`
	// The 2 digit country code associated with the billing ID
	Country string `json:"country"`
}
