/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DomainReturn struct {
	Name      string      `json:"name"`
	Summary   interface{} `json:"summary,omitempty"`
	DirectsTo string      `json:"directsTo"`
	DomainId  int32       `json:"domainId"`
}
