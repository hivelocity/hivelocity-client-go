/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BillingInfo struct {
	// The first name of the individual associated with the billing ID
	First string `json:"first,omitempty"`
	// The bank account number
	AchAccount string `json:"achAccount,omitempty"`
	// The state associated with the billing ID
	State string `json:"state,omitempty"`
	// The zip/postal code associated with the billing ID
	Zip string `json:"zip,omitempty"`
	// The name of the bank associated with the billing ID
	AchBank string `json:"achBank,omitempty"`
	// The email address associated with the billing ID
	Email string `json:"email,omitempty"`
	// The company name associated with the billing ID
	Company string `json:"company,omitempty"`
	// The bank account type (Checking or Savings)
	AchType string `json:"achType,omitempty"`
	// The street address associated with the billing ID
	Address string `json:"address,omitempty"`
	// The verification status of the Credit Card in question
	IsCCVerified bool `json:"isCCVerified,omitempty"`
	// Yes or No for the Pending Verification status
	PendingVerification bool `json:"pendingVerification,omitempty"`
	// The unique Billing ID used to associate Billing Info
	Id int32 `json:"id,omitempty"`
	// The last name of the individual associated with the billing ID
	Last string `json:"last,omitempty"`
	// True if this is the default payment method for future orders
	Default_ bool `json:"default,omitempty"`
	// The country associated with the billing ID
	Country string `json:"country,omitempty"`
	// The expiration date for the credit card
	CcExpire string `json:"ccExpire,omitempty"`
	// The active status associated with the billing ID (0/1)
	Status int32 `json:"status,omitempty"`
	// The credit card network associated with the billing ID
	CcType string `json:"ccType,omitempty"`
	// The city associated with the billing ID in question
	City string `json:"city,omitempty"`
	// The credit card account number
	CcNum string `json:"ccNum,omitempty"`
	// The identification number for the client associated with the Billing info ID
	ClientId int32 `json:"clientId,omitempty"`
	// The bank routing number
	AchAba string `json:"achAba,omitempty"`
	// The type of payment associated with the billing ID ie: CC, ACH etc
	PaymentType string `json:"paymentType,omitempty"`
	// The phone number associated with the billing ID
	Phone string `json:"phone,omitempty"`
}
