/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// BareMetalDeviceUpdate struct for BareMetalDeviceUpdate
type BareMetalDeviceUpdate struct {
	PublicSshKeyId int32    `json:"publicSshKeyId,omitempty"`
	OsName         string   `json:"osName"`
	Script         string   `json:"script,omitempty"`
	Hostname       string   `json:"hostname"`
	Tags           []string `json:"tags,omitempty"`
}