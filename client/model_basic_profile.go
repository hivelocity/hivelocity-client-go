/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// BasicProfile struct for BasicProfile
type BasicProfile struct {
	Last     string                 `json:"last,omitempty"`
	State    map[string]interface{} `json:"state,omitempty"`
	MetaData map[string]interface{} `json:"metaData,omitempty"`
	Id       int32                  `json:"id,omitempty"`
	IsClient bool                   `json:"isClient,omitempty"`
	First    string                 `json:"first,omitempty"`
	Login    string                 `json:"login,omitempty"`
	Country  map[string]interface{} `json:"country,omitempty"`
	Company  map[string]interface{} `json:"company,omitempty"`
	Zip      map[string]interface{} `json:"zip,omitempty"`
	FullName map[string]interface{} `json:"fullName,omitempty"`
	Email    string                 `json:"email,omitempty"`
	City     map[string]interface{} `json:"city,omitempty"`
}
