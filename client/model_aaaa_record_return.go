/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type AaaaRecordReturn struct {
	Address  string `json:"address,omitempty"`
	Id       int32  `json:"id,omitempty"`
	Ttl      int32  `json:"ttl"`
	Type_    string `json:"type"`
	Name     string `json:"name"`
	DomainId int32  `json:"domainId"`
}
