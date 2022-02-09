/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client

// OrderGroup struct for OrderGroup
type OrderGroup struct {
	Name      string  `json:"name,omitempty"`
	SameRack  bool    `json:"same_rack,omitempty"`
	DeviceIds []int32 `json:"device_ids,omitempty"`
	Id        int32   `json:"id,omitempty"`
}
