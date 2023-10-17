/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type PreviewEffectiveIgnition struct {
	PrimaryInterface *PrimaryInterface `json:"primaryInterface,omitempty"`
	Bonded           bool              `json:"bonded,omitempty"`
	PublicKeys       []string          `json:"publicKeys,omitempty"`
	Password         string            `json:"password,omitempty"`
	Hostname         string            `json:"hostname,omitempty"`
	Ignition         interface{}       `json:"ignition"`
}
