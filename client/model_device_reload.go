/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// DeviceReload struct for DeviceReload
type DeviceReload struct {
	Script            string  `json:"script,omitempty"`
	ControlPanelId    int32   `json:"controlPanelId,omitempty"`
	Body              string  `json:"body,omitempty"`
	OperatingSystemId int32   `json:"operatingSystemId"`
	PublicSshKeyIds   []int32 `json:"publicSshKeyIds,omitempty"`
}