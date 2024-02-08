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
	Phone       string `json:"phone,omitempty"`
	Active      int32  `json:"active"`
	Email       string `json:"email"`
	Password    string `json:"password,omitempty"`
	FullName    string `json:"fullName"`
	Description string `json:"description,omitempty"`
}
