/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PreviewEffectiveIgnitionResponse struct {
	Result        interface{}    `json:"result,omitempty"`
	Original      interface{}    `json:"original,omitempty"`
	Modifications []Modification `json:"modifications,omitempty"`
}
