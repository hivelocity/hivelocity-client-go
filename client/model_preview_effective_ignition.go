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
	Password         string            `json:"password,omitempty"`
	Ignition         interface{}       `json:"ignition"`
	Bonded           bool              `json:"bonded,omitempty"`
	PublicKeys       []string          `json:"publicKeys,omitempty"`
	PrimaryInterface *PrimaryInterface `json:"primaryInterface,omitempty"`
	Hostname         string            `json:"hostname,omitempty"`
}