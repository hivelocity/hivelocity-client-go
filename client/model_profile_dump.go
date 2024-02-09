/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ProfileDump struct {
	Login    string      `json:"login,omitempty"`
	First    string      `json:"first,omitempty"`
	Address  interface{} `json:"address,omitempty"`
	Phone    string      `json:"phone,omitempty"`
	Id       int32       `json:"id,omitempty"`
	State    interface{} `json:"state,omitempty"`
	Zip      interface{} `json:"zip,omitempty"`
	Fax      interface{} `json:"fax,omitempty"`
	Email    string      `json:"email,omitempty"`
	Company  interface{} `json:"company,omitempty"`
	Last     string      `json:"last,omitempty"`
	City     interface{} `json:"city,omitempty"`
	MetaData interface{} `json:"metaData,omitempty"`
	FullName interface{} `json:"fullName,omitempty"`
	Created  interface{} `json:"created,omitempty"`
	IsClient bool        `json:"isClient,omitempty"`
	Country  interface{} `json:"country,omitempty"`
}