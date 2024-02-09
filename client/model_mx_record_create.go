/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type MxRecordCreate struct {
	Name       string `json:"name"`
	Ttl        int32  `json:"ttl"`
	Exchange   string `json:"exchange"`
	Preference int32  `json:"preference"`
}
