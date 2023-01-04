/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ContactCreate struct {
	Description string `json:"description,omitempty"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	Active      int32  `json:"active"`
	FullName    string `json:"fullName"`
	Phone       string `json:"phone,omitempty"`
}
