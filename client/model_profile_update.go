/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProfileUpdate struct {
	Address                 string `json:"address,omitempty"`
	First                   string `json:"first,omitempty"`
	DirectToZendesk         string `json:"directToZendesk,omitempty"`
	State                   string `json:"state,omitempty"`
	Zip                     string `json:"zip,omitempty"`
	Email                   string `json:"email,omitempty"`
	Company                 string `json:"company,omitempty"`
	Last                    string `json:"last,omitempty"`
	Password                string `json:"password,omitempty"`
	City                    string `json:"city,omitempty"`
	RequireOauthLogin       string `json:"requireOauthLogin,omitempty"`
	FullName                string `json:"fullName,omitempty"`
	RequirePayPalOauthLogin string `json:"requirePayPalOauthLogin,omitempty"`
	DisableInvoiceEmails    bool   `json:"disableInvoiceEmails,omitempty"`
	Phone                   string `json:"phone,omitempty"`
	Country                 string `json:"country,omitempty"`
}
