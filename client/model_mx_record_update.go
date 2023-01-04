/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MxRecordUpdate struct {
	Exchange   string `json:"exchange"`
	Preference int32  `json:"preference"`
	Name       string `json:"name,omitempty"`
	Ttl        int32  `json:"ttl,omitempty"`
}
