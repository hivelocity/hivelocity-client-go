/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// ProfileUpdate struct for ProfileUpdate
type ProfileUpdate struct {
	Password                string `json:"password,omitempty"`
	RequireOauthLogin       string `json:"requireOauthLogin,omitempty"`
	Phone                   string `json:"phone,omitempty"`
	Last                    string `json:"last,omitempty"`
	State                   string `json:"state,omitempty"`
	First                   string `json:"first,omitempty"`
	Address                 string `json:"address,omitempty"`
	Country                 string `json:"country,omitempty"`
	Company                 string `json:"company,omitempty"`
	Zip                     string `json:"zip,omitempty"`
	FullName                string `json:"fullName,omitempty"`
	Email                   string `json:"email,omitempty"`
	City                    string `json:"city,omitempty"`
	RequirePayPalOauthLogin string `json:"requirePayPalOauthLogin,omitempty"`
}
